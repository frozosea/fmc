package init

import (
	"database/sql"
	"fmt"
	"freight_service/docs"
	"freight_service/internal/city"
	"freight_service/internal/company"
	"freight_service/internal/container"
	"freight_service/internal/freight"
	"freight_service/pkg/cache"
	"freight_service/pkg/logging"
	"freight_service/pkg/middleware"
	pb "freight_service/pkg/proto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"net"
	"os"
	"strings"
	"time"
)

type (
	RedisSettings struct {
		Url string
		Ttl string
	}
	DataBaseSettings struct {
		Host             string
		DatabaseUser     string
		DatabasePassword string
		Database         string
	}
	CityLogsConfig struct {
		CityLogsSaveDir string
	}
	ContactLogsConfig struct {
		ContactLogsSaveDir string
	}
	FreightLogsConfig struct {
		FreightLogsSaveDir string
	}
	ContainerLogsConfig struct {
		SaveDir string
	}
)

func SetupDatabaseConfig() *DataBaseSettings {
	DbSettings := new(DataBaseSettings)
	DbSettings.DatabaseUser = os.Getenv(`POSTGRES_USER`)
	DbSettings.DatabasePassword = os.Getenv(`POSTGRES_PASSWORD`)
	DbSettings.Database = os.Getenv(`POSTGRES_DATABASE`)
	DbSettings.Host = os.Getenv("POSTGRES_HOST")
	return DbSettings
}
func GetDatabase() (*sql.DB, error) {
	dbConf := SetupDatabaseConfig()
	db, err := sql.Open(`pgx`, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConf.Host,
		5432,
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
func getCacheSettings() (*RedisSettings, error) {
	settings := new(RedisSettings)
	cfg, err := ini.Load(`conf/config.ini`)
	sectionRead := cfg.Section("CACHE_SETTINGS")
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if err := sectionRead.MapTo(&settings); err != nil {
		return settings, nil
	}
	settings.Url = os.Getenv("REDIS_URL")
	return settings, nil
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
func GetCache(redisConf *RedisSettings) cache.ICache {
	return cache.NewCache(redis.NewClient(&redis.Options{
		Addr:     redisConf.Url,
		Password: "", // no password set
		DB:       0,  // use default DB
	}), parseExpiration(redisConf.Ttl))
}
func getCityLogsConfig() *CityLogsConfig {
	settings := new(CityLogsConfig)
	cfg, err := ini.Load(`conf/config.ini`)
	sectionRead := cfg.Section("CITY_LOGS")
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if err := sectionRead.MapTo(&settings); err != nil {
		return settings
	}
	return settings
}
func getContactLogsConfig() *ContactLogsConfig {
	settings := new(ContactLogsConfig)
	cfg, err := ini.Load(`conf/config.ini`)
	sectionRead := cfg.Section("CONTACT_LOGS")
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if err := sectionRead.MapTo(&settings); err != nil {
		return settings
	}
	return settings
}
func getFreightLogsConfig() *FreightLogsConfig {
	settings := new(FreightLogsConfig)
	cfg, err := ini.Load(`conf/config.ini`)
	sectionRead := cfg.Section("FREIGHT_LOGS")
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if err := sectionRead.MapTo(&settings); err != nil {
		return settings
	}
	return settings
}

func getContainerLoggerSettings() *ContainerLogsConfig {
	settings := new(ContainerLogsConfig)
	cfg, err := ini.Load(`conf/config.ini`)
	sectionRead := cfg.Section("CONTAINERS_LOGS")
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if err := sectionRead.MapTo(&settings); err != nil {
		return settings
	}
	return settings
}

var redisSettings, _ = getCacheSettings()
var RedisCache = GetCache(redisSettings)
var DataBase, _ = GetDatabase()

var CityRepository = city.NewRepository(DataBase)
var CityLoggerConfig = getCityLogsConfig()
var CityLogger = logging.NewLogger(CityLoggerConfig.CityLogsSaveDir)
var CityService = city.NewService(CityRepository, CityLogger, RedisCache)
var CityGrpcHandler = city.NewGrpc(CityService)
var CityHttpHandler = city.NewHttp(CityService)

var CompanyRepository = company.NewRepository(DataBase)
var CompanyLoggerConfig = getContactLogsConfig()
var CompanyLogger = logging.NewLogger(CompanyLoggerConfig.ContactLogsSaveDir)
var CompanyService = company.NewService(CompanyRepository, CompanyLogger, RedisCache)
var CompanyGrpcHandler = company.NewGrpc(CompanyService)
var CompanyHttpHandler = company.NewHttp(CompanyService)

var FreightRepository = freight.NewRepository(DataBase)
var FreightLoggerConfig = getFreightLogsConfig()
var FreightLogger = logging.NewLogger(FreightLoggerConfig.FreightLogsSaveDir)
var FreightService = freight.NewService(FreightRepository, FreightLogger, RedisCache)
var FreightGrpcHandler = freight.NewGrpc(FreightService)
var FreightHttpHandler = freight.NewHttp(FreightService)

var ContainerRepository = container.NewRepository(DataBase)
var ContainerLoggerConfig = getContainerLoggerSettings()
var ContainerLogger = logging.NewLogger(ContainerLoggerConfig.SaveDir)
var ContainerService = container.NewService(ContainerLogger, ContainerRepository)
var ContainerHttpHandler = container.NewHttp(ContainerService)
var ContainerGrpcHandler = container.NewGrpc(ContainerService)

func ConfigureAndRunGrpc() error {
	server := grpc.NewServer()
	pb.RegisterCityServiceServer(server, CityGrpcHandler)
	pb.RegisterCompanyServiceServer(server, CompanyGrpcHandler)
	pb.RegisterFreightServiceServer(server, FreightGrpcHandler)
	pb.RegisterContainersServiceServer(server, ContainerGrpcHandler)
	l, err := net.Listen("tcp", fmt.Sprintf(`:%d`, 51839))
	if err != nil {
		return err
	}
	return server.Serve(l)
}

func ConfigureAndRunHttp() error {
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(middleware.New().CheckAuth)
	{
		r.GET("/freights", FreightHttpHandler.GetFreights)
		r.POST("/freight", FreightHttpHandler.AddFreight)
		r.PUT("/freight", FreightHttpHandler.UpdateFreight)
		r.DELETE("/freight", FreightHttpHandler.DeleteFreight)
	}
	{
		r.GET("/countries", CityHttpHandler.GetAllCountries)
		r.GET("/cities", CityHttpHandler.GetAllCities)
		r.POST("/country", CityHttpHandler.AddCountry)
		r.POST("/city", CityHttpHandler.AddCity)
		r.PUT("/country", CityHttpHandler.UpdateCountry)
		r.PUT("/city", CityHttpHandler.UpdateCity)
		r.DELETE("/country", CityHttpHandler.DeleteCountry)
		r.DELETE("/city", CityHttpHandler.DeleteCity)
	}
	{
		r.GET("/companies", CompanyHttpHandler.GetAll)
		r.POST("/company", CompanyHttpHandler.Add)
		r.PUT("/company", CompanyHttpHandler.UpdateCompany)
		r.DELETE("/company", CompanyHttpHandler.Delete)
	}
	{
		r.GET("/containers", ContainerHttpHandler.GetAllContainers)
		r.POST("/container", ContainerHttpHandler.AddContainer)
		r.PUT("/container", ContainerHttpHandler.UpdateContainer)
		r.DELETE("/container", ContainerHttpHandler.DeleteContainer)
	}
	{
		docs.SwaggerInfo.BasePath = "/"
		r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	return r.Run("0.0.0.0:8090")
}
