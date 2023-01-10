package main

import (
	"context"
	scheduleTrackingPb "github.com/frozosea/fmc-pb/schedule-tracking"
	trackingPb "github.com/frozosea/fmc-pb/tracking"
	userPb "github.com/frozosea/fmc-pb/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/alts"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Builder struct {
	variables *EnvVariables
	mux       *runtime.ServeMux
	ctx       context.Context
	opts      []grpc.DialOption
}

func NewBuilder() *Builder {
	return &Builder{
		ctx:  context.Background(),
		mux:  runtime.NewServeMux(),
		opts: []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())},
	}
}

func (b *Builder) initEnvVariables() *Builder {
	v, err := getEnvVariables()
	if err != nil {
		panic(err)
		return nil
	}
	b.variables = v
	return b
}

func (b *Builder) initUserGateway() *Builder {
	clientOpts := alts.DefaultClientOptions()
	clientOpts.TargetServiceAccounts = []string{b.variables.AltsKeyForUserApp}
	altsTC := alts.NewClientCreds(clientOpts)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(altsTC)}
	url := getUrl(b.variables.UserAppIp, b.variables.UserAppPort)
	if err := userPb.RegisterUserFeedbackHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	if err := userPb.RegisterUserHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	if err := userPb.RegisterAuthHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	return b
}
func (b *Builder) initTrackingGateway() *Builder {
	clientOpts := alts.DefaultClientOptions()
	clientOpts.TargetServiceAccounts = []string{b.variables.AltsKeyForTrackingApp}
	altsTC := alts.NewClientCreds(clientOpts)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(altsTC)}
	url := getUrl(b.variables.TrackingAppIp, b.variables.TrackingAppPort)
	if err := trackingPb.RegisterTrackingByContainerNumberHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	if err := trackingPb.RegisterTrackingByBillNumberHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	if err := trackingPb.RegisterScacServiceHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	return b
}
func (b *Builder) initScheduleTrackingGateway() *Builder {
	clientOpts := alts.DefaultClientOptions()
	clientOpts.TargetServiceAccounts = []string{b.variables.AltsKeyForScheduleTrackingApp}
	altsTC := alts.NewClientCreds(clientOpts)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(altsTC)}
	url := getUrl(b.variables.ScheduleTrackingIp, b.variables.ScheduleTrackingPort)
	if err := scheduleTrackingPb.RegisterScheduleTrackingHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	if err := scheduleTrackingPb.RegisterArchiveHandlerFromEndpoint(b.ctx, b.mux, url, opts); err != nil {
		panic(err)
		return nil
	}
	return b
}

func (b *Builder) Run() {
	go func() {
		log.Println("START HTTP GATEWAY SERVER")
		if err := http.ListenAndServe("0.0.0.0:8080", b.mux); err != nil {
			panic(err)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-s
}
