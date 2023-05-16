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
	"log"
	"os"
	"schedule-tracking/internal/archive"
	"schedule-tracking/internal/domain"
	excel_writer "schedule-tracking/pkg/excel-writer"
	"schedule-tracking/pkg/logging"
	"schedule-tracking/pkg/scheduler"
	"schedule-tracking/pkg/tracking"
	"schedule-tracking/pkg/util"
	"strconv"
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
		Host     string
		Port     int
		Email    string
		Password string
	}
	TrackingClientSettings struct {
		Ip   string
		Port string
	}
	UserClientSettings struct {
		Ip   string
		Port string
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
	ArchiveLoggerSettings struct {
		SaveDir string
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
	emailSender.Email = os.Getenv("SENDER_EMAIL")
	emailSender.Password = os.Getenv("EMAIL_PASSWORD")
	emailSender.Host = os.Getenv("EMAIL_SMTP_HOST")
	port, err := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
	if err != nil {
		panic(err)
	}
	emailSender.Port = port
	return emailSender
}
func GetTrackingSettings() (*TrackingClientSettings, error) {
	clientSettings := new(TrackingClientSettings)
	ip, port := os.Getenv("TRACKING_GRPC_HOST"), os.Getenv("TRACKING_GRPC_PORT")
	if ip == "" || port == "" {
		return nil, errors.New("no env variables")
	}
	clientSettings.Ip = ip
	clientSettings.Port = port
	return clientSettings, nil
}

func GetTrackingClient(conf *TrackingClientSettings, logger logging.ILogger) *tracking.Client {
	var conn *grpc.ClientConn
	var err error
	if os.Getenv("PRODUCTION") == "1" {
		tlsCredentials, err := loadClientTLSCredentials()
		if err != nil {
			panic(err)
		}

		conn, err = grpc.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), grpc.WithTransportCredentials(tlsCredentials))
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
	}
	client := tracking.NewClient(conn, logger)
	return client
}
func GetUserScheduleTrackingClient(conf *UserClientSettings, logger logging.ILogger) *domain.UserClient {
	var conn *grpc.ClientConn
	var err error
	if os.Getenv("PRODUCTION") == "1" {
		tlsCredentials, err := loadClientTLSCredentials()
		if err != nil {
			panic(err)
		}

		conn, err = grpc.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), grpc.WithTransportCredentials(tlsCredentials))
		if err != nil {
			panic(err)
		}
	} else {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%s", conf.Ip, conf.Port), grpc.WithInsecure())
		if err != nil {
			panic(err)
		}
	}
	var pbClient = user_pb.NewScheduleTrackingClient(conn)
	return domain.NewClient(pbClient, logger)
}
func getScheduleTrackingLoggingConfig() (*ScheduleTrackingLoggerSettings, error) {
	config := new(ScheduleTrackingLoggerSettings)
	return readIni("SCHEDULE_LOGS", config)
}
func GetTimeFormatterSettings() (*TimeFormatterSettings, error) {
	e := os.Getenv("TIME_FORMAT")
	if e == "" {
		return nil, errors.New("NO TIME_FORMAT ENV VARIABLE")
	}
	return &TimeFormatterSettings{Format: e}, nil
}
func GetArchiveLoggerSettings() (*ArchiveLoggerSettings, error) {
	config := new(ArchiveLoggerSettings)
	return readIni("ARCHIVE_LOGS", config)
}

func GetUserClientSettings() (*UserClientSettings, error) {
	ip, port := os.Getenv("USER_GRPC_HOST"), os.Getenv("USER_GRPC_PORT")
	settings := new(UserClientSettings)
	if ip == "" || port == "" {
		return nil, errors.New("no env variables")
	}
	settings.Ip, settings.Port = ip, port
	return settings, nil
}

func GetScheduleTrackingAndArchiveGrpcService() *domain.Grpc {
	trackerConf, err := GetTrackingSettings()
	if err != nil {
		panic(err)
	}

	loggerConf, err := getScheduleTrackingLoggingConfig()
	if err != nil {
		panic(err)
	}
	userConf, err := GetUserClientSettings()
	if err != nil {
		panic(err)
	}

	arrivedChecker := tracking.NewArrivedChecker()
	controllerLogger := logging.NewLogger(loggerConf.ControllerSaveDir)
	excelWriter := excel_writer.NewWriter(os.Getenv("PWD"))
	sender := GetEmailSenderSettings()
	emailSender := mailing.NewMailing(sender.Host, sender.Port, sender.Email, sender.Password)
	format, err := GetTimeFormatterSettings()
	if err != nil {
		panic(err)
	}
	client := GetUserScheduleTrackingClient(userConf, logging.NewLogger(loggerConf.ClientSaveDir))
	timeFormatter := domain.NewTimeFormatter(format.Format)
	db, err := GetDatabase()
	if err != nil {
		panic(err)
	}
	repository := domain.NewRepository(db)
	archiveRepository := archive.NewRepository(db)
	archiveLoggerSettings, err := GetArchiveLoggerSettings()
	if err != nil {
		return nil
	}
	var timezone = os.Getenv("TZ")
	if timezone == "" {
		timezone = "Asia/Vladivostok"
	}
	var taskManager = scheduler.NewDefault(timezone)
	archiveService := archive.NewService(logging.NewLogger(archiveLoggerSettings.SaveDir), archiveRepository)

	taskGetter := domain.NewCustomTasks(
		GetTrackingClient(trackerConf, logging.NewLogger(loggerConf.TrackingResultSaveDir)),
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

	scheduleTrackingService := domain.NewService(controllerLogger, client, taskManager, ExcelTrackingResultSaveDir, repository, taskGetter)
	if recoveryTaskErr := RecoveryTasks(repository, scheduleTrackingService); recoveryTaskErr != nil {
		log.Println(err)
	}
	var conn *grpc.ClientConn
	if os.Getenv("PRODUCTION") == "1" {
		tlsCredentials, err := loadClientTLSCredentials()
		if err != nil {
			panic(err)
		}

		conn, err = grpc.Dial(fmt.Sprintf("%s:%s", userConf.Ip, userConf.Port), grpc.WithTransportCredentials(tlsCredentials))
	} else {
		conn, err = grpc.Dial(fmt.Sprintf("%s:%s", userConf.Ip, userConf.Port), grpc.WithInsecure())
	}
	if err != nil {
		panic(err)
	}
	return domain.NewGrpc(scheduleTrackingService, logging.NewLogger(loggerConf.ServiceSaveDir), util.NewTokenManager(user_pb.NewAuthClient(conn)))
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
