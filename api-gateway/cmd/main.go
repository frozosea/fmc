package main

import (
	init_package "fmc-with-git/internal/init"
	"github.com/joho/godotenv"
	"log"
)

// @title FindMyCargo API
// @version 1.0.0
// @description API server for application
// @host 0.0.0.0:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env file")
	}
	init_package.Run()
}
