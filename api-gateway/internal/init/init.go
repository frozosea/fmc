package init_package

import (
	"errors"
	"fmc-gateway/docs"
	"fmc-gateway/internal/logging"
	"fmc-gateway/internal/middleware"
	"fmc-gateway/internal/utils"
	"fmc-gateway/pkg/auth"
	schedule_tracking "fmc-gateway/pkg/schedule-tracking"
	"fmc-gateway/pkg/tracking"
	"fmc-gateway/pkg/user"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

type (
	AuthSettings struct {
		Ip   string
		Port string
	}
	TrackingSettings struct {
		Ip   string
		Port string
	}
	UserSettings struct {
		Ip   string
		Port string
	}
	ScheduleTrackingSettings struct {
		Ip   string
		Port string
	}
)

func getAuthSettings() (*AuthSettings, error) {
	authSettings := AuthSettings{}
	authSettings.Ip = os.Getenv("USER_APP_IP")
	authSettings.Port = os.Getenv("USER_APP_PORT")
	return &authSettings, nil
}
func getTrackingSettings() (*TrackingSettings, error) {
	trackingSettings := new(TrackingSettings)
	trackingSettings.Ip = os.Getenv("TRACKING_IP")
	trackingSettings.Port = os.Getenv("TRACKING_PORT")
	return trackingSettings, nil

}
func getScheduleTrackingSettings() (*ScheduleTrackingSettings, error) {
	settings := new(ScheduleTrackingSettings)
	ip, port := os.Getenv("SCHEDULE_TRACKING_HOST"), os.Getenv("SCHEDULE_TRACKING_PORT")
	if ip == "" || port == "" {
		return nil, errors.New("no env variables")
	}
	settings.Ip, settings.Port = ip, port
	return settings, nil
}
func GetTrackingClient(ip, port string, logger logging.ILogger) (*tracking.Client, error) {
	var url string
	if ip == "" {
		url = fmt.Sprintf(`localhost:%s`, port)
	}
	url = fmt.Sprintf(`%s:%s`, ip, port)
	trackingConnection, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return tracking.NewClient(trackingConnection, logger), nil
}

func getTrackingHttpHandler(client *tracking.Client, utils *utils.HttpUtils) *tracking.HttpHandler {
	return tracking.NewHttpHandler(client, utils)
}

func initTrackingRoutes(router *gin.Engine, handler *tracking.HttpHandler, middleware *middleware.Middleware) {
	trackingGroup := router.Group(`/tracking`)
	trackingGroup.Use(middleware.CheckAccessMiddleware)
	{
		trackingGroup.POST(`/trackByBillNumber`, handler.TrackByBillNumber)
		trackingGroup.POST(`/trackByContainerNumber`, handler.TrackByContainerNumber)
	}
}
func getScheduleTrackingClient(ip, port string, logger logging.ILogger) (*schedule_tracking.Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return schedule_tracking.NewClient(conn, logger), nil
}
func getScheduleTrackingHttpHandler(client *schedule_tracking.Client, utils *utils.HttpUtils) *schedule_tracking.HttpHandler {
	return schedule_tracking.NewHttpHandler(client, utils)
}
func initScheduleRoutes(router *gin.Engine, handler *schedule_tracking.HttpHandler, middleware *middleware.Middleware) {
	group := router.Group(`/schedule`)
	group.Use(middleware.CheckAccessMiddleware)
	{
		group.POST(`/addContainer`, handler.AddContainersOnTrack)
		group.POST(`/addBillNo`, handler.AddBillNumbersOnTrack)
		group.PUT(`/updateTime`, handler.UpdateTrackingTime)
		group.PUT(`/addEmail`, handler.AddEmailsOnTracking)
		group.DELETE(`/deleteEmail`, handler.DeleteEmailFromTrack)
		group.DELETE(`/deleteContainers`, handler.DeleteContainersFromTrack)
		group.DELETE(`/deleteBillNumbers`, handler.DeleteBillNumbersFromTrack)
		group.POST(`/getInfo`, handler.GetInfoAboutTracking)
	}
	router.GET(`/schedule/getTz`, handler.GetTimeZone)

}

func getAuthClient(ip, port string, logger logging.ILogger) (*auth.Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
		return &auth.Client{}, err
	}
	return auth.NewClient(conn, logger), nil
}
func getAuthHttpHandler(client *auth.Client) *auth.HttpHandler {
	return auth.NewHttpHandler(client)
}
func initAuthRoutes(router *gin.Engine, AuthHttpHandler *auth.HttpHandler) {
	group := router.Group(`/auth`)
	{
		group.POST(`/refresh`, AuthHttpHandler.Refresh)
		router.POST(`/auth/register`, AuthHttpHandler.Register)
		router.POST(`/auth/login`, AuthHttpHandler.Login)
	}
}
func getUserClient(ip, port string, logger logging.ILogger) (*user.Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &user.Client{}, err
	}
	return user.NewClient(conn, logger), nil
}

func getUserHttpHandler(client *user.Client, utils *utils.HttpUtils) *user.HttpHandler {
	return user.NewHttpHandler(client, utils)
}
func initUserRoutes(router *gin.Engine, handler *user.HttpHandler, middleware *middleware.Middleware) {
	group := router.Group(`/user`)
	group.Use(middleware.CheckAccessMiddleware)
	{
		group.POST(`/addContainers`, handler.AddContainersToAccount)
		group.POST(`/addBillNumbers`, handler.AddBillNumbersToAccount)
		group.DELETE(`/deleteContainers`, handler.DeleteContainersFromAccount)
		group.DELETE(`/deleteBillNumbers`, handler.DeleteBillNumbersFromAccount)
		group.GET(`/getAllBillsContainers`, handler.GetAll)
	}
}
func initDocsRoutes(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Run() {
	var AuthServerSettings, err = getAuthSettings()
	if err != nil {
		panic(err)
		return
	}
	var AuthLogger = logging.NewLogger("authLogs")
	var AuthClient, exc = getAuthClient(AuthServerSettings.Ip, AuthServerSettings.Port, AuthLogger)
	if exc != nil {
		panic(exc)
		return
	}
	var AuthHttpHandler = getAuthHttpHandler(AuthClient)

	var httpUtils = utils.NewHttpUtils(AuthClient)

	var TrackingServerSettings, getTrackingSettingsErr = getTrackingSettings()
	if getTrackingSettingsErr != nil {
		panic(getTrackingSettingsErr)
		return
	}
	var TrackingClientLogger = logging.NewLogger("trackingLogs")
	var TrackingClient, getTrackingClientErr = GetTrackingClient(TrackingServerSettings.Ip, TrackingServerSettings.Port, TrackingClientLogger)
	if getTrackingClientErr != nil {
		panic(getTrackingClientErr)
		return
	}
	var TrackingHttpHandler = getTrackingHttpHandler(TrackingClient, httpUtils)

	var UserLogger = logging.NewLogger("userLogs")
	var UserClient, getUserClientErr = getUserClient(AuthServerSettings.Ip, AuthServerSettings.Port, UserLogger)
	if getUserClientErr != nil {
		panic(getUserClientErr)
		return
	}
	var UserHttpHandler = getUserHttpHandler(UserClient, httpUtils)

	var scheduleTrackingSettings, getSettingsErr = getScheduleTrackingSettings()
	if getSettingsErr != nil {
		panic(getSettingsErr)
		return
	}
	var ScheduleTrackingLogger = logging.NewLogger("scheduleTrackingLogs")
	var ScheduleTrackingClient, getScheduleTrackingClientErr = getScheduleTrackingClient(scheduleTrackingSettings.Ip, scheduleTrackingSettings.Port, ScheduleTrackingLogger)
	if getScheduleTrackingClientErr != nil {
		panic(getScheduleTrackingClientErr)
		return
	}
	var ScheduleTrackingHttpHandler = getScheduleTrackingHttpHandler(ScheduleTrackingClient, httpUtils)

	var Middleware = middleware.NewMiddleware(httpUtils, AuthClient)

	router := gin.Default()
	initAuthRoutes(router, AuthHttpHandler)
	initTrackingRoutes(router, TrackingHttpHandler, Middleware)
	initDocsRoutes(router)
	initUserRoutes(router, UserHttpHandler, Middleware)
	initScheduleRoutes(router, ScheduleTrackingHttpHandler, Middleware)
	defaultCors := cors.DefaultConfig()
	defaultCors.AllowAllOrigins = true
	router.Use(cors.New(defaultCors))
	log.Fatal(router.Run(fmt.Sprintf(`0.0.0.0:8080`)))
}
