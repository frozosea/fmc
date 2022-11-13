package init_package

import (
	"errors"
	"fmc-gateway/docs"
	"fmc-gateway/internal/auth"
	freight_service "fmc-gateway/internal/freight-service"
	"fmc-gateway/internal/history"
	"fmc-gateway/internal/schedule-tracking"
	tracking "fmc-gateway/internal/tracking"
	"fmc-gateway/internal/user"
	"fmc-gateway/pkg/logging"
	"fmc-gateway/pkg/middleware"
	"fmc-gateway/pkg/utils"
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
	FreightClientSettings struct {
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
func getFreightClientSettings() (*FreightClientSettings, error) {
	settings := new(FreightClientSettings)
	//TODO add these env variables
	ip, port := os.Getenv("FREIGHT_HOST"), os.Getenv("FREIGHT_PORT")
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
func GetScheduleTrackingHistoryClient(ip, port string) (*history.ScheduleTrackingTasksClient, error) {
	var url string
	if ip == "" {
		url = fmt.Sprintf(`localhost:%s`, port)
	}
	url = fmt.Sprintf(`%s:%s`, ip, port)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return history.NewScheduleTrackingTasksClient(conn), nil
}
func getFreightClient(ip, port string) (freight_service.IClient, error) {
	var url string
	if ip == "" {
		url = fmt.Sprintf(`localhost:%s`, port)
	}
	url = fmt.Sprintf(`%s:%s`, ip, port)
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return freight_service.NewClient(conn), nil
}
func getTrackingHttpHandler(client *tracking.Client, utils *utils.HttpUtils) *tracking.HttpHandler {
	return tracking.NewHttpHandler(client, utils)
}

func initTrackingRoutes(router *gin.Engine, handler *tracking.HttpHandler) {
	trackingGroup := router.Group(`/tracking`)
	{
		trackingGroup.GET(`/bill`, handler.TrackByBillNumber)
		trackingGroup.GET(`/container`, handler.TrackByContainerNumber)
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
		group.POST(`/containers`, handler.AddContainersOnTrack)
		group.POST(`/bills`, handler.AddBillNumbersOnTrack)
		group.PUT(`/bills`, handler.UpdateBills)
		group.PUT(`/containers`, handler.UpdateContainers)
		group.DELETE(`/containers`, handler.DeleteContainersFromTrack)
		group.DELETE(`/bills`, handler.DeleteBillNumbersFromTrack)
		group.GET(`/info`, handler.GetInfoAboutTracking)
	}
	router.GET(`/schedule/timezone`, handler.GetTimeZone)

}
func initHistoryRoutes(router *gin.Engine, handler *history.HttpHandler, middleware *middleware.Middleware) {
	group := router.Group(`/history`)
	group.Use(middleware.CheckAccessMiddleware)
	{
		group.GET("/tasks", handler.GetTasksArchive)
	}
}
func getAuthClient(ip, port string, logger logging.ILogger) (*auth.Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
		return &auth.Client{}, err
	}
	return auth.NewClient(conn, logger), nil
}
func getAuthHttpHandler(client auth.IClient) *auth.HttpHandler {
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
		group.POST(`/containers`, handler.AddContainersToAccount)
		group.POST(`/bills`, handler.AddBillNumbersToAccount)
		group.DELETE(`/containers`, handler.DeleteContainersFromAccount)
		group.DELETE(`/bills`, handler.DeleteBillNumbersFromAccount)
		group.GET(`/all`, handler.GetAll)
	}
}
func initFreightsRouter(router *gin.Engine, handler *freight_service.Http) {
	group := router.Group("/freight")
	{
		group.GET("/freights", handler.GetFreights)
		group.GET("/cities", handler.GetAllCities)
		group.GET("/companies", handler.GetAllCompanies)
		group.GET("/containers", handler.GetAllContainers)
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
	scheduleTrackingHistoryClient, err := GetScheduleTrackingHistoryClient(scheduleTrackingSettings.Ip, scheduleTrackingSettings.Port)
	if err != nil {
		panic(getSettingsErr)
		return
	}
	//freightClientSettings, err := getFreightClientSettings()
	//if err != nil {
	//	panic(getSettingsErr)
	//	return
	//}
	//freightClient, err := getFreightClient(freightClientSettings.Ip, freightClientSettings.Port)
	//if err != nil {
	//	panic(getSettingsErr)
	//	return
	//}
	router := gin.Default()
	initAuthRoutes(router, AuthHttpHandler)
	initTrackingRoutes(router, TrackingHttpHandler)
	initDocsRoutes(router)
	initUserRoutes(router, UserHttpHandler, Middleware)
	initScheduleRoutes(router, ScheduleTrackingHttpHandler, Middleware)
	initHistoryRoutes(router, history.NewHttpHandler(scheduleTrackingHistoryClient, httpUtils), Middleware)
	//initFreightsRouter(router, freight_service.NewHttp(freightClient))
	defaultCors := cors.DefaultConfig()
	defaultCors.AddAllowMethods("GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS")
	defaultCors.AllowHeaders = []string{"Accept", "Authorization", "Cache-Control", "Content-Type", "DNT", "If-Modified-Since", "Keep-Alive", "Origin", "User-Agent", "X-Requested-With", "X-Real-Ip"}
	defaultCors.AllowAllOrigins = true
	defaultCors.AllowCredentials = true
	router.Use(cors.New(defaultCors))
	log.Fatal(router.Run(fmt.Sprintf(`0.0.0.0:8080`)))
}
