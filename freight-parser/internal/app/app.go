package app

import (
	"database/sql"
	"fmc-newest/internal/cache"
	"fmc-newest/internal/logging"
	"fmc-newest/pkg/city"
	"fmc-newest/pkg/contact"
	"fmc-newest/pkg/freight"
	"fmc-newest/pkg/line"
	pb "fmc-newest/pkg/proto"
	"fmt"
	"github.com/go-ini/ini"
	"github.com/go-redis/redis/v8"
	_ "github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
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
	ServerSettings struct {
		Port string
	}
	DataBaseSettings struct {
		Host             string
		DatabaseUser     string
		DatabasePassword string
		Database         string
	}
	LineLogsConfig struct {
		LineLogsSaveDir string
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
)

func readIni[T comparable](section string, settingsModel *T) *T {
	cfg, err := ini.Load(`conf/config.ini`)
	sectionRead := cfg.Section(section)
	if err != nil {
		log.Fatalf(`read config from ini file err:%s`, err)
	}
	if err := sectionRead.MapTo(&settingsModel); err != nil {
		return settingsModel
	}
	return settingsModel
}
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
func GetServerSettings() *ServerSettings {
	config := new(ServerSettings)
	config.Port = os.Getenv("GRPC_PORT")
	return config
}
func getCacheSettings() (*RedisSettings, error) {
	settings := new(RedisSettings)
	return readIni("CACHE_SETTINGS", settings), nil
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
	cfg := new(CityLogsConfig)
	return readIni("CITY_LOGS", cfg)
}
func getContactLogsConfig() *ContactLogsConfig {
	cfg := new(ContactLogsConfig)
	return readIni("CONTACT_LOGS", cfg)
}
func getFreightLogsConfig() *FreightLogsConfig {
	cfg := new(FreightLogsConfig)
	return readIni("FREIGHT_LOGS", cfg)
}
func getLineLogsConfig() *LineLogsConfig {
	settings := new(LineLogsConfig)
	return readIni("LINE_LOGS", settings)
}

var redisSettings, _ = getCacheSettings()
var RedisCache = GetCache(redisSettings)
var DataBase, _ = GetDatabase()

var CityRepository = city.NewRepository(DataBase)
var CityLoggerConfig = getCityLogsConfig()
var CityLogger = logging.NewLogger(CityLoggerConfig.CityLogsSaveDir)
var CityController = city.NewController(CityRepository, CityLogger, RedisCache)
var CityService = city.NewService(CityController)

var ContactRepository = contact.NewRepository(DataBase)
var ContactLoggerConfig = getContactLogsConfig()
var ContactLogger = logging.NewLogger(ContactLoggerConfig.ContactLogsSaveDir)
var ContactController = contact.NewController(ContactRepository, ContactLogger, RedisCache)
var ContactService = contact.NewService(ContactController)

var FreightRepository = freight.NewRepository(DataBase)
var FreightLoggerConfig = getFreightLogsConfig()
var FreightLogger = logging.NewLogger(FreightLoggerConfig.FreightLogsSaveDir)
var FreightController = freight.NewController(FreightRepository, FreightLogger, RedisCache)
var FreightService = freight.NewGetFreightService(FreightController, FreightLogger)

var LineRepository = line.NewRepository(DataBase)
var LineLoggerConfig = getLineLogsConfig()
var LineLogger = logging.NewLogger(LineLoggerConfig.LineLogsSaveDir)
var LineController = line.NewController(LineRepository, LineLogger, RedisCache)
var LineService = line.NewService(LineController)

func ConfigureAndRun() error {
	server := grpc.NewServer()
	pb.RegisterCityServiceServer(server, CityService)
	pb.RegisterContactServiceServer(server, ContactService)
	pb.RegisterFreightServiceServer(server, FreightService)
	pb.RegisterLineServiceServer(server, LineService)
	settings := GetServerSettings()
	l, err := net.Listen("tcp", fmt.Sprintf(`:%s`, settings.Port))
	if err != nil {
		return err
	}
	return server.Serve(l)
}
