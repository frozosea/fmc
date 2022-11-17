package initpackage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"user-api/internal/auth"
	"user-api/internal/feedback"
	"user-api/internal/schedule-tracking"
	"user-api/internal/user"
	"user-api/pkg/cache"
	"user-api/pkg/logging"
	"user-api/pkg/mailing"
)

type (
	DataBaseSettings struct {
		Host             string
		DatabaseUser     string
		DatabasePassword string
		Database         string
		Port             string
	}
	JwtSettings struct {
		AccessTokenExpiration  string
		RefreshTokenExpiration string
		JwtSecretKey           string
	}
	RedisSettings struct {
		Url string
		Ttl string
	}
	ScheduleTrackingLoggerSettings struct {
		ServiceSaveDir string
	}
	UserLoggerSettings struct {
		ControllerSaveDir string
	}
	AuthLoggerSettings struct {
		ControllerSaveDir string
		ServiceSaveDir    string
	}
	MailingSettings struct {
		Host         string
		Port         int
		Email        string
		Password     string
		SendToEmails []string
	}
)

func readIni[T comparable](section string, settingsModel *T) (*T, error) {
	cfg, err := ini.Load(`conf/cfg.ini`)
	sectionRead := cfg.Section(section)
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
		return settingsModel, err
	}
	if readErr := sectionRead.MapTo(&settingsModel); readErr != nil {
		return settingsModel, readErr
	}
	return settingsModel, nil
}
func SetupDatabaseConfig() *DataBaseSettings {
	DbSettings := new(DataBaseSettings)
	DbSettings.DatabaseUser = os.Getenv(`POSTGRES_USER`)
	DbSettings.DatabasePassword = os.Getenv(`POSTGRES_PASSWORD`)
	DbSettings.Database = os.Getenv(`POSTGRES_DATABASE`)
	DbSettings.Host = os.Getenv("POSTGRES_HOST")
	DbSettings.Port = os.Getenv("POSTGRES_PORT")
	return DbSettings
}

func GetDatabase() (*sql.DB, error) {
	dbConf := SetupDatabaseConfig()
	db, err := sql.Open(`pgx`, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		dbConf.Port,
		dbConf.DatabaseUser,
		dbConf.DatabasePassword,
		dbConf.Database))
	if err != nil {
		log.Fatalf(`open database err:%s`, err.Error())
		return db, err
	}
	if exc := db.Ping(); exc != nil {
		return db, exc
	}
	return db, nil
}

func GetJwtSecret() string {
	return os.Getenv(`JWT_SECRET_KEY`)
}

func GetTokenSettings() (*JwtSettings, error) {
	jwt := new(JwtSettings)
	jwt.JwtSecretKey = GetJwtSecret()
	return readIni("TOKEN", jwt)
}

func GetRedisSettings() *RedisSettings {
	redisCli := new(RedisSettings)
	redisCli.Url = os.Getenv("REDIS_URL")
	redisCli.Ttl = os.Getenv("REDIS_TTL")
	return redisCli
}
func getScheduleTrackingLoggingConfig() (*ScheduleTrackingLoggerSettings, error) {
	config := new(ScheduleTrackingLoggerSettings)
	return readIni("SCHEDULE_LOGS", config)
}
func GetScheduleTrackingService(db *sql.DB) *schedule_tracking.Service {
	loggerConf, err := getScheduleTrackingLoggingConfig()
	if err != nil {
		panic(err)
		return nil
	}
	repository := schedule_tracking.NewRepository(db)
	return schedule_tracking.NewService(repository, logging.NewLogger(loggerConf.ServiceSaveDir), redisCache)
}
func parseTime(timeStr string, sep string) int64 {
	splitInfo := strings.Split(timeStr, sep)
	var exp int64
	if _, err := fmt.Sscanf(splitInfo[0], `%d`, &exp); err != nil {
		panic(err)
	}
	return exp
}

func parseExpiration(parseString string) time.Duration {
	if strings.Contains(parseString, "h") {
		return time.Duration(parseTime(parseString, "h")) * time.Hour
	} else if strings.Contains(parseString, "m") {
		return time.Duration(parseTime(parseString, "m")) * time.Minute
	} else {
		return time.Second
	}
}
func getAuthLoggerConfig() (*AuthLoggerSettings, error) {
	config := new(AuthLoggerSettings)
	return readIni("AUTH_LOGS", config)
}

func GetAuthGrpcService(db *sql.DB) *auth.Grpc {
	loggerConf, err := getAuthLoggerConfig()
	if err != nil {
		panic(err)
	}
	tokenSettings, getTokenSettingsErr := GetTokenSettings()
	if getTokenSettingsErr != nil {
		panic(getTokenSettingsErr)
	}
	tokenManager := auth.NewTokenManager(tokenSettings.JwtSecretKey, parseExpiration(tokenSettings.AccessTokenExpiration), parseExpiration(tokenSettings.RefreshTokenExpiration))
	hash := auth.NewHash()
	repository := auth.NewRepository(db, hash)
	mSettings, err := getMailingSettings()
	if err != nil {
		panic(err)
		return nil
	}
	m := mailing.NewMailing(mSettings.Host, mSettings.Port, mSettings.Email, mSettings.Password)
	controller := auth.NewService(repository, tokenManager, hash, m, logging.NewLogger(loggerConf.ControllerSaveDir))
	return auth.NewGrpc(controller, logging.NewLogger(loggerConf.ServiceSaveDir))
}
func getUserLoggerConfig() (*UserLoggerSettings, error) {
	config := new(UserLoggerSettings)
	return readIni("USER_LOGS", config)
}
func GetCache(redisConf *RedisSettings) cache.ICache {
	redisCache := cache.NewCache(redis.NewClient(&redis.Options{
		Addr:     redisConf.Url,
		Password: "", // no password set
		DB:       0,  // use default DB
	}), parseExpiration(redisConf.Ttl))
	return redisCache
}

var redisConf = GetRedisSettings()
var redisCache = GetCache(redisConf)

func GetUserGrpcService(db *sql.DB, redisConf *RedisSettings) *user.Grpc {
	loggerConf, err := getUserLoggerConfig()
	if err != nil {
		panic(err)
	}
	repository := user.NewRepository(db)
	controller := user.NewService(repository, logging.NewLogger(loggerConf.ControllerSaveDir), redisCache)
	return user.NewGrpc(controller)
}
func getMailingSettings() (*MailingSettings, error) {
	email, password, smtpHost, smtpPort, sendToEmails := os.Getenv("SEND_EMAIL"), os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"), os.Getenv("SEND_TO_EMAILS")
	if email == "" || password == "" || smtpHost == "" || smtpPort == "" || sendToEmails == "" {
		return nil, errors.New("no env variables")
	}
	intPort, err := strconv.Atoi(smtpPort)
	if err != nil {
		return nil, err
	}
	var sendToEmailArr []string
	for _, v := range strings.Split(sendToEmails, ";") {
		sendToEmailArr = append(sendToEmailArr, v)
	}
	return &MailingSettings{
		Host:         smtpHost,
		Port:         intPort,
		Email:        email,
		Password:     password,
		SendToEmails: sendToEmailArr,
	}, nil
}
func GetFeedbackDeliveries(db *sql.DB) (*feedback.Grpc, *feedback.Http) {
	repository := feedback.NewRepository(db)
	mSettings, err := getMailingSettings()
	if err != nil {
		panic(err)
		return nil, nil
	}
	logger := logging.NewLogger("feedback")
	m := mailing.NewMailing(mSettings.Host, mSettings.Port, mSettings.Email, mSettings.Password)
	service := feedback.NewService(m, repository, logger, mSettings.SendToEmails)
	return feedback.NewGrpc(service), feedback.NewHttp(service)
}
