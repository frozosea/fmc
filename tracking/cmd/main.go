package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	_ "net/http/pprof"
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
		initTimeInspector().
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

	go func() {
		panic(http.ListenAndServe("0.0.0.0:9999", nil))
	}()
}
