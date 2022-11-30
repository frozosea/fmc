package main

import (
	"fmt"
	pb "github.com/frozosea/fmc-pb/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user-api/internal/feedback"
	conf "user-api/internal/init"
)

func GetServer() (*grpc.Server, *feedback.Http, error) {
	server := grpc.NewServer()
	db, err := conf.GetDatabase()
	if err != nil {
		return server, nil, err
	}
	feedbackGrpcService, feedbackHttpHandler := conf.GetFeedbackDeliveries(db)
	pb.RegisterAuthServer(server, conf.GetAuthGrpcService(db))
	pb.RegisterUserServer(server, conf.GetUserGrpcService(db, conf.GetRedisSettings()))
	pb.RegisterScheduleTrackingServer(server, conf.GetScheduleTrackingService(db))
	pb.RegisterUserFeedbackServer(server, feedbackGrpcService)
	return server, feedbackHttpHandler, nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("no env file")
	}
	server, feedbackHttpHandler, err := GetServer()
	if err != nil {
		panic(err)
		return
	}
	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf(`0.0.0.0:9001`))
		if err != nil {
			panic(err)
			return
		}
		log.Println("START GRPC SERVER")
		if err := server.Serve(l); err != nil {
			panic(err)
		}
	}()
	go func() {
		r := gin.Default()
		r.Use(feedback.NewMiddleware().CheckAdminAccess)
		{
			r.GET("/all", feedbackHttpHandler.GetAll)
			r.GET("/byEmail", feedbackHttpHandler.GetByEmail)
		}
		fmt.Println("START HTTP")
		log.Fatal(r.Run("0.0.0.0:9005"))
	}()
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-s
}
