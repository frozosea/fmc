package main

import (
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file")
	}

	NewBuilder().
		initEnvVariables().
		initGRPCServer().
		initDatabase().
		initSitcStore().
		initCaptchaSolver().
		initUnlocodesRepo().
		initCache().
		initContainerTrackers().
		initMainContainerTracker().
		initScacRepository().
		initLoginProvider().
		initContainerTrackingService().
		initContainerTrackingGRPCService().
		registerContainerTrackingGRPCService().
		initBillTrackers().
		initBillMainTracker().
		initBillTrackingService().
		initBillTrackingGRPCService().
		registerBillTrackingGRPCService().
		initScacService().
		initScacGrpcService().
		registerScacGrpcService().
		Run()
}
