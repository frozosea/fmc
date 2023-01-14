package main

import (
	"context"
	scheduleTrackingPb "github.com/frozosea/fmc-pb/schedule-tracking"
	trackingPb "github.com/frozosea/fmc-pb/tracking"
	userPb "github.com/frozosea/fmc-pb/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
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
	builder := &Builder{
		ctx:  context.Background(),
		mux:  runtime.NewServeMux(),
		opts: []grpc.DialOption{},
	}
	builder.initTlsCreds()
	return builder
}
func (b *Builder) initTlsCreds() *Builder {
	b.opts = append(b.opts, grpc.WithInsecure())
	return b
}
func (b *Builder) initEnvVariables() *Builder {
	v, err := getEnvVariables()
	if err != nil {
		panic(err)
	}
	b.variables = v
	return b
}

func (b *Builder) initUserGateway() *Builder {
	url := getUrl(b.variables.UserAppIp, b.variables.UserAppPort)
	if err := userPb.RegisterUserFeedbackHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	if err := userPb.RegisterUserHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	if err := userPb.RegisterAuthHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	return b
}
func (b *Builder) initTrackingGateway() *Builder {
	url := getUrl(b.variables.TrackingAppIp, b.variables.TrackingAppPort)
	if err := trackingPb.RegisterTrackingByContainerNumberHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	if err := trackingPb.RegisterTrackingByBillNumberHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	if err := trackingPb.RegisterScacServiceHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	return b
}
func (b *Builder) initScheduleTrackingGateway() *Builder {
	url := getUrl(b.variables.ScheduleTrackingIp, b.variables.ScheduleTrackingPort)
	if err := scheduleTrackingPb.RegisterScheduleTrackingHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
	}
	if err := scheduleTrackingPb.RegisterArchiveHandlerFromEndpoint(b.ctx, b.mux, url, b.opts); err != nil {
		panic(err)
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
