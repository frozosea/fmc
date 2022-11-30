package main

import (
	"fmt"
	pb "github.com/frozosea/fmc-pb/schedule-tracking"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	init_package "schedule-tracking/internal/init"
	"syscall"
)

func Run() {
	s := grpc.NewServer()
	scheduleTrackingGrpcService := init_package.GetScheduleTrackingAndArchiveGrpcService()
	pb.RegisterScheduleTrackingServer(s, scheduleTrackingGrpcService)
	l, err := net.Listen("tcp", `0.0.0.0:8005`)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println("START ON 0.0.0.0:8005")
	log.Fatal(s.Serve(l))
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env file")
	}

	go func() {
		Run()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
