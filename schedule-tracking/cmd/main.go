package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	init_package "schedule-tracking/internal/init"
	pb "schedule-tracking/pkg/proto"
	"syscall"
)

func Run() {
	s := grpc.NewServer()
	service := init_package.GetScheduleTrackingService()
	pb.RegisterScheduleTrackingServer(s, service)
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
