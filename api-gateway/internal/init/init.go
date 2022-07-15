package init_package

import (
	"fmc-with-git/docs"
	"fmc-with-git/internal/logging"
	"fmc-with-git/internal/middleware"
	"fmc-with-git/internal/utils"
	"fmc-with-git/pkg/auth"
	schedule_tracking "fmc-with-git/pkg/schedule-tracking"
	"fmc-with-git/pkg/tracking"
	"fmc-with-git/pkg/user"
	"fmt"
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
)

func getAuthSettings() (*AuthSettings, error) {
	authSettings := AuthSettings{}
	authSettings.Ip = os.Getenv("AUTH_IP")
	authSettings.Port = os.Getenv("AUTH_PORT")
	return &authSettings, nil
}
func getTrackingSettings() (*TrackingSettings, error) {
	trackingSettings := new(TrackingSettings)
	trackingSettings.Ip = os.Getenv("TRACKING_IP")
	trackingSettings.Port = os.Getenv("TRACKING_PORT")
	return trackingSettings, nil

}
func GetTrackingClient(ip, port string, logger logging.ILogger) (*tracking.Client, error) {
	var url string
	if ip == "" {
		url = fmt.Sprintf(`:%s`, port)
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

func initTrackingRoutes(router *gin.Engine, utils *utils.HttpUtils, middleware *middleware.Middleware, ip, port string, logger logging.ILogger) {
	client, err := GetTrackingClient(ip, port, logger)
	if err != nil {
		panic(err)
		return
	}
	handler := getTrackingHttpHandler(client, utils)
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
func initScheduleRoutes(router *gin.Engine, utils *utils.HttpUtils, middleware *middleware.Middleware, ip, port string, logger logging.ILogger) {
	client, err := getScheduleTrackingClient(ip, port, logger)
	if err != nil {
		panic(err)
	}
	handler := getScheduleTrackingHttpHandler(client, utils)
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
}

func getAuthClient(ip, port string, logger logging.ILogger) (*auth.Client, error) {
	conn, err := grpc.Dial(fmt.Sprintf(`%s:%s`, ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &auth.Client{}, err
	}
	return auth.NewClient(conn, logger), nil
}
func getAuthHttpHandler(client *auth.Client, utils *utils.HttpUtils) *auth.HttpHandler {
	return auth.NewHttpHandler(client, utils)
}
func initAuthRoutes(router *gin.Engine, utils *utils.HttpUtils, ip, port string, logger logging.ILogger) {
	client, err := getAuthClient(ip, port, logger)
	if err != nil {
		panic(err)
	}
	handler := getAuthHttpHandler(client, utils)
	group := router.Group(`/auth`)
	{
		group.POST(`/refresh`, handler.Refresh)
		router.POST(`/auth/register`, handler.Register)
		router.POST(`/auth/login`, handler.Login)
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
func initUserRoutes(router *gin.Engine, utils *utils.HttpUtils, middleware *middleware.Middleware, ip, port string, logger logging.ILogger) {
	client, err := getUserClient(ip, port, logger)
	if err != nil {
		panic(err.Error())
	}
	handler := getUserHttpHandler(client, utils)
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
	router := gin.Default()
	httpUtils := utils.NewHttpUtils(os.Getenv("JWT_SECRET_KEY"))
	authMiddleware := middleware.NewMiddleware(httpUtils)
	authSettings, err := getAuthSettings()
	if err != nil {
		panic(err)
		return
	}
	trackingSettings, err := getTrackingSettings()
	if err != nil {
		panic(err)
	}
	fmt.Println(trackingSettings)
	initAuthRoutes(router, httpUtils, authSettings.Ip, authSettings.Port, logging.NewLogger("authLogs"))
	initTrackingRoutes(router, httpUtils, authMiddleware, trackingSettings.Ip, trackingSettings.Port, logging.NewLogger("trackingLogs"))
	initDocsRoutes(router)
	initUserRoutes(router, httpUtils, authMiddleware, authSettings.Ip, authSettings.Port, logging.NewLogger("userLogs"))
	initScheduleRoutes(router, httpUtils, authMiddleware, authSettings.Ip, authSettings.Port, logging.NewLogger("schedule"))
	log.Fatal(router.Run(fmt.Sprintf(`0.0.0.0:8080`)))
}
