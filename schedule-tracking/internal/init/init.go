package initpackage

import (
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	excel_writer "schedule-tracking/internal/excel-writer"
	"schedule-tracking/internal/logging"
	"schedule-tracking/internal/mailing"
	"schedule-tracking/internal/scheduler"
	"schedule-tracking/internal/tracking"
	user_pb "schedule-tracking/internal/user-pb"
	"schedule-tracking/pkg/domain"
)

type (
	EmailSenderSettings struct {
		SenderName      string
		SenderEmail     string
		UnisenderApiKey string
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
	EmailSignatureSettings struct {
		EmailSignature string
	}
)

const ExcelTrackingResultSaveDir = "."

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
	ip, port := os.Getenv("TRACKING_GRPC_HOST"), os.Getenv("TRACKING_GRPC_PORT")
	if ip == "" || port == "" {
		return nil, errors.New("no env variables")
	}
	clientSettings.Ip = ip
	clientSettings.Port = port
	return clientSettings, nil
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
func GetUserScheduleTrackingClient(conf *UserClientSettings, logger logging.ILogger) *domain.Client {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, conf.Ip, conf.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
		return &domain.Client{}
	}
	var pbClient = user_pb.NewScheduleTrackingClient(conn)
	return domain.NewClient(pbClient, logger)
}
func getScheduleTrackingLoggingConfig() (*ScheduleTrackingLoggerSettings, error) {
	config := new(ScheduleTrackingLoggerSettings)
	return readIni("SCHEDULE_LOGS", config)
}
func GetTimeFormatterSettings() (*TimeFormatterSettings, error) {
	config := new(TimeFormatterSettings)
	return readIni("TIME_FORMAT", config)
}

var TaskManager = scheduler.NewDefault()

func GetUserClientSettings() (*UserClientSettings, error) {
	ip, port := os.Getenv("USER_GRPC_HOST"), os.Getenv("USER_GRPC_PORT")
	settings := new(UserClientSettings)
	if ip == "" || port == "" {
		return nil, errors.New("no env variables")
	}
	settings.Ip, settings.Port = ip, port
	return settings, nil
}

func GetScheduleTrackingService() *domain.Service {
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
	client := GetUserScheduleTrackingClient(userConf, logging.NewLogger(loggerConf.ClientSaveDir))
	timeFormatter := domain.NewTimeFormatter(format.Format)
	taskGetter := domain.NewCustomTasks(GetTrackingClient(trackerConf, logging.NewLogger(loggerConf.TrackingResultSaveDir)), client, arrivedChecker, logging.NewLogger(loggerConf.TaskGetterSaveDir), excelWriter, emailSender, timeFormatter)
	controller := domain.NewController(controllerLogger, client, TaskManager, ExcelTrackingResultSaveDir, taskGetter)
	return domain.NewService(controller, logging.NewLogger(loggerConf.ServiceSaveDir))
}
