package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"net"
	conf "user-api/internal/init"
	pb "user-api/pkg/proto"
)

func GetServer() (*grpc.Server, error) {
	server := grpc.NewServer()
	db, err := conf.GetDatabase()
	if err != nil {
		return server, err
	}
	pb.RegisterAuthServer(server, conf.GetAuthService(db))
	pb.RegisterUserServer(server, conf.GetUserService(db, conf.GetRedisSettings()))
	pb.RegisterScheduleTrackingServer(server, conf.GetScheduleTrackingService(db))
	fmt.Println("start scheduler")
	//go conf.TaskManager.Start()
	return server, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env file")
	}
	server, err := GetServer()
	if err != nil {
		panic(err)
		return
	}
	l, err := net.Listen("tcp", fmt.Sprintf(`0.0.0.0:51472`))
	if err != nil {
		panic(err)
		return
	}
	if err := server.Serve(l); err != nil {
		panic(err)
	}
}
