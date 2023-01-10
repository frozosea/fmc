package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pb "github.com/frozosea/fmc-pb/user"
	"github.com/frozosea/mailing"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
	"user-api/internal/auth"
	"user-api/internal/feedback"
	"user-api/internal/schedule-tracking"
	"user-api/internal/user"
	"user-api/pkg/cache"
	"user-api/pkg/logging"
	"user-api/pkg/util"
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
		AuthKey      string
	}
	AuthSettings struct {
		AltsKey string
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
	jwt.AccessTokenExpiration = os.Getenv("ACCESS_TOKEN_EXPIRATION")
	if jwt.AccessTokenExpiration == "" {
		return nil, errors.New("no ACCESS_TOKEN_EXPIRATION env variable")
	}
	jwt.RefreshTokenExpiration = os.Getenv("REFRESH_TOKEN_EXPIRATION")
	if jwt.RefreshTokenExpiration == "" {
		return nil, errors.New("no REFRESH_TOKEN_EXPIRATION env variable")
	}
	return jwt, nil
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
func GetScheduleTrackingService(db *sql.DB) *schedule_tracking.Grpc {
	loggerConf, err := getScheduleTrackingLoggingConfig()
	if err != nil {
		panic(err)
		return nil
	}
	repository := schedule_tracking.NewRepository(db)
	return schedule_tracking.NewGrpc(repository, logging.NewLogger(loggerConf.ServiceSaveDir), redisCache)
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

func getJwtTokenManager() (auth.ITokenManager, error) {
	tokenSettings, err := GetTokenSettings()
	if err != nil {
		return nil, err
	}
	tokenManager := auth.NewTokenManager(tokenSettings.JwtSecretKey, parseExpiration(tokenSettings.AccessTokenExpiration), parseExpiration(tokenSettings.RefreshTokenExpiration))
	return tokenManager, nil
}

func GetAuthGrpcService(db *sql.DB) *auth.Grpc {
	loggerConf, err := getAuthLoggerConfig()
	if err != nil {
		panic(err)
	}
	tokenManager, err := getJwtTokenManager()
	if err != nil {
		panic(err)
		return nil
	}
	hash := auth.NewHash()
	repository := auth.NewRepository(db, hash)
	mSettings, err := getMailingSettings()
	if err != nil {
		panic(err)
		return nil
	}
	m, err := mailing.NewWithElasticEmail(mSettings.Host, mSettings.Port, mSettings.Email, mSettings.Password, mSettings.AuthKey, "containerTrackingContacts")
	if err != nil {
		panic(err)
		return nil
	}
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

func getAuthSettings() (*AuthSettings, error) {
	key := os.Getenv("ALTS_KEY")
	if key == "" {
		return nil, errors.New("no env variable")
	}
	return &AuthSettings{AltsKey: key}, nil
}

func GetUserGrpcService(db *sql.DB) *user.Grpc {
	jwtManager, err := getJwtTokenManager()
	if err != nil {
		panic(err)
		return nil
	}
	loggerConf, err := getUserLoggerConfig()
	if err != nil {
		panic(err)
	}
	scheduleTrackingMicroserviceIp, scheduleTrackingMicroservicePort := os.Getenv("SCHEDULE_TRACKING_IP"), os.Getenv("SCHEDULE_TRACKING_PORT")
	if scheduleTrackingMicroserviceIp == "" || scheduleTrackingMicroservicePort == "" {
		panic("no env variables!")
		return nil
	}
	url := fmt.Sprintf("%s:%s", scheduleTrackingMicroserviceIp, scheduleTrackingMicroservicePort)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
		return nil
	}
	scheduleTrackingInfoRepository := user.NewScheduleTrackingInfoRepository(conn)
	repository := user.NewRepository(db, scheduleTrackingInfoRepository)
	service := user.NewService(repository, logging.NewLogger(loggerConf.ControllerSaveDir), redisCache)
	authSettings, err := getAuthSettings()
	if err != nil {
		panic(err)
		return nil
	}
	return user.NewGrpc(service, util.NewTokenManager(jwtManager), authSettings.AltsKey)
}

func getMailingSettings() (*MailingSettings, error) {
	email, password, smtpHost, smtpPort, sendToEmails, authKey := os.Getenv("SEND_EMAIL"), os.Getenv("PASSWORD"), os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"), os.Getenv("SEND_TO_EMAILS"), os.Getenv("AUTH_KEY")
	if email == "" || password == "" || smtpHost == "" || smtpPort == "" || sendToEmails == "" || authKey == "" {
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
		AuthKey:      authKey,
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
	m, err := mailing.NewWithElasticEmail(mSettings.Host, mSettings.Port, mSettings.Email, mSettings.Password, mSettings.AuthKey, "containerTrackingContacts")
	if err != nil {
		panic(err)
		return nil, nil
	}
	service := feedback.NewService(m, repository, logger, mSettings.SendToEmails)
	return feedback.NewGrpc(service), feedback.NewHttp(service)
}

func GetServer() (*grpc.Server, *feedback.Http, error) {
	authSettings, err := getAuthSettings()
	if err != nil {
		return nil, nil, err
	}
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	server := grpc.NewServer(
		grpc.Creds(altsTC),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
				if err := alts.ClientAuthorizationCheck(ctx, []string{authSettings.AltsKey}); err != nil {
					return ctx, status.Error(codes.Unauthenticated, err.Error())
				}
				return ctx, nil
			}),
		)),
	)
	db, err := GetDatabase()
	if err != nil {
		return nil, nil, err
	}
	feedbackGrpcService, feedbackHttpHandler := GetFeedbackDeliveries(db)
	pb.RegisterAuthServer(server, GetAuthGrpcService(db))
	pb.RegisterUserServer(server, GetUserGrpcService(db))
	pb.RegisterScheduleTrackingServer(server, GetScheduleTrackingService(db))
	pb.RegisterUserFeedbackServer(server, feedbackGrpcService)
	return server, feedbackHttpHandler, nil
}

func BuildAndRun() {
	server, feedbackHttpHandler, err := GetServer()
	if err != nil {
		panic(err)
		return
	}
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf(`0.0.0.0:9001`))
		if err != nil {
			panic(err)
			return
		}
		log.Println("START GRPC SERVER")
		if err := server.Serve(l); err != nil {
			panic(err)
		}
	}()
	go func() {
		r := gin.Default()
		r.Use(feedback.NewMiddleware().CheckAdminAccess)
		{
			r.GET("/all", feedbackHttpHandler.GetAll)
			r.GET("/byEmail", feedbackHttpHandler.GetByEmail)
		}
		fmt.Println("START HTTP")
		log.Fatal(r.Run("0.0.0.0:9005"))
	}()
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-s
}
