package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	user_pb "github.com/frozosea/fmc-pb/user"
	"github.com/frozosea/mailing"
	"github.com/go-ini/ini"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"log"
	"os"
	"schedule-tracking/internal/archive"
	"schedule-tracking/internal/domain"
	excel_writer "schedule-tracking/pkg/excel-writer"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/scheduler"
	"schedule-tracking/pkg/tracking"
	"schedule-tracking/pkg/util"
)

type (
	DataBaseSettings struct {
		Host             string
		DatabaseUser     string
		DatabasePassword string
		Database         string
		Port             string
	}
	EmailSenderSettings struct {
		SenderName      string
		SenderEmail     string
		UnisenderApiKey string
	}
	TrackingClientSettings struct {
		Ip      string
		Port    string
		AltsKey string
	}
	UserClientSettings struct {
		Ip      string
		Port    string
		AltsKey string
	}
	TimeFormatterSettings struct {
		Format string
	}
	ScheduleTrackingLoggerSettings struct {
		ClientSaveDir         string
		TrackingResultSaveDir string
		ServiceSaveDir        string
		ControllerSaveDir     string
		TaskGetterSaveDir     string
	}
	EmailSignatureSettings struct {
		EmailSignature string
	}
	ArchiveLoggerSettings struct {
		SaveDir string
	}
	AuthSettings struct {
		AltsKey string
	}
)

const ExcelTrackingResultSaveDir = "."

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
	db, err := sql.Open(`postgres`, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
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

func GetEmailSenderSettings() *EmailSenderSettings {
	emailSender := new(EmailSenderSettings)
	emailSender.SenderEmail = os.Getenv("SENDER_EMAIL")
	emailSender.SenderName = os.Getenv("SENDER_NAME")
	emailSender.UnisenderApiKey = os.Getenv("UNISENDER_API_KEY")
	return emailSender
}

func GetTrackingSettings() (*TrackingClientSettings, error) {
	clientSettings := new(TrackingClientSettings)
	ip, port, altsKey := os.Getenv("TRACKING_GRPC_HOST"), os.Getenv("TRACKING_GRPC_PORT"), os.Getenv("ALTS_KEY_FOR_TRACKING_API")
	if ip == "" || port == "" || altsKey == "" {
		return nil, errors.New("no env variables")
	}
	clientSettings.Ip = ip
	clientSettings.Port = port
	clientSettings.AltsKey = altsKey
	return clientSettings, nil
}

func GetEmailSignature() (*EmailSignatureSettings, error) {
	signature := os.Getenv("EMAIL_MESSAGE_SIGNATURE")
	if signature == "" {
		return nil, errors.New("no env variable")
	}
	return &EmailSignatureSettings{EmailSignature: signature}, nil
}

func GetMailing(sender *EmailSenderSettings) mailing.IMailing {
	settings, err := GetEmailSignature()
	if err != nil {
		panic(err)
	}
	return mailing.NewWithUniSender(sender.SenderName, sender.SenderEmail, sender.UnisenderApiKey, settings.EmailSignature)
}

func GetTrackingClient(conf *TrackingClientSettings, logger logging.ILogger, altsKey string) *tracking.Client {
	clientOpts := alts.DefaultClientOptions()
	clientOpts.TargetServiceAccounts = []string{altsKey}
	altsTC := alts.NewClientCreds(clientOpts)
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, conf.Ip, conf.Port), grpc.WithTransportCredentials(altsTC))
	if err != nil {
		panic(err)
		return &tracking.Client{}
	}
	client := tracking.NewClient(conn, logger)
	return client
}

func GetUserScheduleTrackingClient(conf *UserClientSettings, logger logging.ILogger, altsKey string) *domain.UserClient {
	clientOpts := alts.DefaultClientOptions()
	clientOpts.TargetServiceAccounts = []string{altsKey}
	altsTC := alts.NewClientCreds(clientOpts)
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, conf.Ip, conf.Port), grpc.WithTransportCredentials(altsTC))
	if err != nil {
		panic(err)
		return &domain.UserClient{}
	}
	var pbClient = user_pb.NewScheduleTrackingClient(conn)
	return domain.NewClient(pbClient, logger)
}

func GetAuthClient(conf *UserClientSettings) (user_pb.AuthClient, error) {
	clientOpts := alts.DefaultClientOptions()
	clientOpts.TargetServiceAccounts = []string{conf.AltsKey}
	altsTC := alts.NewClientCreds(clientOpts)
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, conf.Ip, conf.Port), grpc.WithTransportCredentials(altsTC))
	if err != nil {
		return nil, err
	}
	return user_pb.NewAuthClient(conn), nil
}

func getScheduleTrackingLoggingConfig() (*ScheduleTrackingLoggerSettings, error) {
	config := new(ScheduleTrackingLoggerSettings)
	return readIni("SCHEDULE_LOGS", config)
}

func GetTimeFormatterSettings() (*TimeFormatterSettings, error) {
	timeFormat := os.Getenv("TIME_FORMAT")
	if timeFormat == "" {
		return nil, errors.New("no env variable")
	}
	return &TimeFormatterSettings{Format: timeFormat}, nil
}

func GetArchiveLoggerSettings() (*ArchiveLoggerSettings, error) {
	config := new(ArchiveLoggerSettings)
	return readIni("ARCHIVE_LOGS", config)
}

func GetUserClientSettings() (*UserClientSettings, error) {
	ip, port, altsKey := os.Getenv("USER_GRPC_HOST"), os.Getenv("USER_GRPC_PORT"), os.Getenv("ALTS_KEY_FOR_USER_API")
	settings := new(UserClientSettings)
	if ip == "" || port == "" || altsKey == "" {
		return nil, errors.New("no env variables")
	}
	settings.Ip, settings.Port, settings.AltsKey = ip, port, altsKey
	return settings, nil
}

func GetScheduleTracking() *domain.Grpc {
	trackerConf, getSettingsErr := GetTrackingSettings()
	if getSettingsErr != nil {
		panic(getSettingsErr)
	}
	loggerConf, err := getScheduleTrackingLoggingConfig()
	if err != nil {
		panic(err)
	}
	userConf, getUserSettingsErr := GetUserClientSettings()
	if getUserSettingsErr != nil {
		panic(getUserSettingsErr)
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
	client := GetUserScheduleTrackingClient(userConf, logging.NewLogger(loggerConf.ClientSaveDir), userConf.AltsKey)
	timeFormatter := domain.NewTimeFormatter(format.Format)
	db, getDbErr := GetDatabase()
	if getDbErr != nil {
		panic(getDbErr)
	}
	repository := domain.NewRepository(db)
	archiveRepository := archive.NewRepository(db)
	archiveLoggerSettings, err := GetArchiveLoggerSettings()
	if err != nil {
		return nil
	}
	archiveService := archive.NewService(logging.NewLogger(archiveLoggerSettings.SaveDir), archiveRepository)
	var timezone = os.Getenv("TZ")
	if timezone == "" {
		timezone = "Asia/Vladivostok"
	}
	var taskManager = scheduler.NewDefault(timezone)
	taskGetter := domain.NewCustomTasks(
		GetTrackingClient(trackerConf, logging.NewLogger(loggerConf.TrackingResultSaveDir), trackerConf.AltsKey),
		client,
		arrivedChecker,
		logging.NewLogger(loggerConf.TaskGetterSaveDir),
		excelWriter,
		emailSender,
		timeFormatter,
		repository,
		archiveService,
		taskManager,
	)
	authClient, err := GetAuthClient(userConf)
	if err != nil {
		panic(err)
		return nil
	}
	scheduleTrackingService := domain.NewService(controllerLogger, client, taskManager, ExcelTrackingResultSaveDir, repository, taskGetter)
	if recoveryTaskErr := RecoveryTasks(repository, scheduleTrackingService); recoveryTaskErr != nil {
		log.Println(err)
	}
	return domain.NewGrpc(scheduleTrackingService, logging.NewLogger(loggerConf.ServiceSaveDir), util.NewTokenManager(authClient))
}

func GetAuthSettings() (*AuthSettings, error) {
	altsKey := os.Getenv("ALTS_KEY")
	if altsKey == "" {
		return nil, errors.New("no env variable")
	}
	return &AuthSettings{AltsKey: altsKey}, nil
}

func RecoveryTasks(repo domain.IRepository, controller *domain.Service) error {
	tasks, err := repo.GetAll(context.Background())
	if err != nil {
		switch err.(type) {
		case *domain.NoTasksError:
			return nil
		default:
			return err
		}
	}
	for _, task := range tasks {
		if !task.IsContainer {
			if _, addErr := controller.AddBillNumbersOnTrack(context.Background(), &domain.BaseTrackReq{
				Numbers:             []string{task.Number},
				UserId:              task.UserId,
				Time:                task.Time,
				Emails:              task.Emails,
				EmailMessageSubject: task.EmailMessageSubject,
			}); addErr != nil {
				return addErr
			}
		} else {
			if _, addErr := controller.AddContainerNumbersOnTrack(context.Background(), &domain.BaseTrackReq{
				Numbers:             []string{task.Number},
				UserId:              task.UserId,
				Time:                task.Time,
				Emails:              task.Emails,
				EmailMessageSubject: task.EmailMessageSubject,
			}); addErr != nil {
				return addErr
			}
		}
	}
	return nil
}
