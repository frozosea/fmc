package main

import (
	"context"
	"fmt"
	pb "github.com/frozosea/fmc-pb/schedule-tracking"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	authSettings, err := GetAuthSettings()
	if err != nil {
		panic(err)
		return
	}
	altsTC := alts.NewServerCreds(alts.DefaultServerOptions())
	s := grpc.NewServer(
		grpc.Creds(altsTC),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
				if err := alts.ClientAuthorizationCheck(ctx, []string{authSettings.AltsKey}); err != nil {
					return ctx, status.Error(codes.Unauthenticated, err.Error())
				}
				return ctx, nil
			}),
		)))
	scheduleTrackingGrpcService := GetScheduleTracking()
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

	go Run()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
