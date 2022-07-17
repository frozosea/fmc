package initpackage

import (
	"database/sql"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
	"time"
	"user-api/internal/cache"
	excel_writer "user-api/internal/excel-writer"
	"user-api/internal/logging"
	"user-api/internal/mailing"
	"user-api/internal/scheduler"
	"user-api/internal/tracking"
	"user-api/pkg/auth"
	schedule_tracking "user-api/pkg/schedule-tracking"
	"user-api/pkg/user"
)

type (
	DataBaseSettings struct {
		Host             string
		DatabaseUser     string
		DatabasePassword string
		Database         string
		Port             string
	}
	ServerSettings struct {
		Port string
	}
	JwtSettings struct {
		AccessTokenExpiration  string
		RefreshTokenExpiration string
		JwtSecretKey           string
	}
	EmailSenderSettings struct {
		SenderName      string
		SenderEmail     string
		UnisenderApiKey string
	}
	RedisSettings struct {
		Url string
		Ttl string
	}
	TrackingClientSettings struct {
		Ip   string
		Port string
	}
	TimeFormatterSettings struct {
		Format string
	}
	ScheduleTrackingLoggerSettings struct {
		TrackingResultSaveDir string
		ServiceSaveDir        string
		ControllerSaveDir     string
		TaskGetterSaveDir     string
	}
	UserLoggerSettings struct {
		ControllerSaveDir string
	}
	AuthLoggerSettings struct {
		ControllerSaveDir string
		ServiceSaveDir    string
	}
	EmailSignatureSettings struct {
		EmailSignature string
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
func GetServerSettings() *ServerSettings {
	config := new(ServerSettings)
	config.Port = os.Getenv("GRPC_PORT")
	return config
}
func GetJwtSecret() string {
	return os.Getenv(`JWT_SECRET_KEY`)
}

func GetTokenSettings() (*JwtSettings, error) {
	jwt := new(JwtSettings)
	jwt.JwtSecretKey = GetJwtSecret()
	return readIni("TOKEN", jwt)
}

func GetEmailSenderSettings() *EmailSenderSettings {
	emailSender := new(EmailSenderSettings)
	emailSender.SenderEmail = os.Getenv("SENDER_EMAIL")
	emailSender.SenderName = os.Getenv("SENDER_NAME")
	emailSender.UnisenderApiKey = os.Getenv("UNISENDER_API_KEY")
	return emailSender
}
func GetRedisSettings() *RedisSettings {
	redisCli := new(RedisSettings)
	redisCli.Url = os.Getenv("REDIS_URL")
	redisCli.Ttl = os.Getenv("REDIS_TTL")
	return redisCli
}
func GetTrackingSettings() (*TrackingClientSettings, error) {
	clientSettings := new(TrackingClientSettings)
	return readIni("TRACKING", clientSettings)
}
func GetEmailSignature() (*EmailSignatureSettings, error) {
	settings := new(EmailSignatureSettings)
	return readIni("EMAIL_SETTINGS", settings)
}
func GetMailing(sender *EmailSenderSettings) *mailing.Mailing {
	logger := logging.NewLogger("emails")
	settings, err := GetEmailSignature()
	if err != nil {
		panic(err)
	}
	return mailing.NewMailing(logger, sender.SenderName, sender.SenderEmail, sender.UnisenderApiKey, settings.EmailSignature)
}
func GetTrackingClient(conf *TrackingClientSettings, logger logging.ILogger) *tracking.Client {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, conf.Ip, conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
		return &tracking.Client{}
	}
	client := tracking.NewClient(conn, logger)
	return client
}
func getScheduleTrackingLoggingConfig() (*ScheduleTrackingLoggerSettings, error) {
	config := new(ScheduleTrackingLoggerSettings)
	return readIni("SCHEDULE_LOGS", config)
}
func GetTimeFormatterSettings() (*TimeFormatterSettings, error) {
	config := new(TimeFormatterSettings)
	return readIni("TIME_FORMAT", config)
}

var TaskManager = scheduler.NewDefaultScheduler()

func GetScheduleTrackingService(db *sql.DB) *schedule_tracking.Service {
	trackerConf, getSettingsErr := GetTrackingSettings()
	if getSettingsErr != nil {
		panic(getSettingsErr)
	}
	loggerConf, err := getScheduleTrackingLoggingConfig()
	if err != nil {
		panic(err)
	}
	arrivedChecker := tracking.NewArrivedChecker()
	controllerLogger := logging.NewLogger(loggerConf.ControllerSaveDir)
	excelWriter := excel_writer.NewWriter(os.Getenv("PWD"))
	sender := GetEmailSenderSettings()
	emailSender := GetMailing(sender)
	format, getTimeFormatErr := GetTimeFormatterSettings()
	if getTimeFormatErr != nil {
		panic(getTimeFormatErr)
	}
	rdsCache := GetCache(GetRedisSettings())
	timeFormatter := schedule_tracking.NewTimeFormatter(format.Format)
	repository := schedule_tracking.NewRepository(db)
	taskGetter := schedule_tracking.NewCustomTasks(GetTrackingClient(trackerConf, logging.NewLogger(loggerConf.TrackingResultSaveDir)), arrivedChecker, logging.NewLogger(loggerConf.TaskGetterSaveDir), excelWriter, emailSender, timeFormatter, repository)
	controller := schedule_tracking.NewController(controllerLogger, repository, TaskManager, taskGetter, rdsCache)
	return schedule_tracking.NewService(controller, logging.NewLogger(loggerConf.ServiceSaveDir))
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

func GetAuthService(db *sql.DB) *auth.Service {
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

	controller := auth.NewController(repository, tokenManager, logging.NewLogger(loggerConf.ControllerSaveDir))
	return auth.NewService(controller, logging.NewLogger(loggerConf.ServiceSaveDir))
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
func GetUserService(db *sql.DB, redisConf *RedisSettings) *user.Service {
	cache := GetCache(redisConf)
	loggerConf, err := getUserLoggerConfig()
	if err != nil {
		panic(err)
	}
	repository := user.NewRepository(db)
	controller := user.NewController(repository, logging.NewLogger(loggerConf.ControllerSaveDir), cache)
	return user.NewService(controller)
}
