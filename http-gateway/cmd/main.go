package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env file")
	}
	NewBuilder().
		initEnvVariables().
		initUserGateway().
		initTrackingGateway().
		initScheduleTrackingGateway().
		Run()
}
