package main

import (
	init_package "freight_service/internal/init"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Freight service API
// @version 1.0.0
// @description API server for admin panel on freight service
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env file")
	}
	go func() {
		log.Println("RUN GRPC SERVER")
		log.Fatal(init_package.ConfigureAndRunGrpc())
	}()
	go func() {
		log.Println("RUN HTTP SERVER FOR ADMIN PANEL")
		log.Fatal(init_package.ConfigureAndRunHttp())
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
