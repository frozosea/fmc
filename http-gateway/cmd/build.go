package main

import (
	"context"
	scheduleTrackingPb "github.com/frozosea/fmc-pb/schedule-tracking"
	trackingPb "github.com/frozosea/fmc-pb/tracking"
	userPb "github.com/frozosea/fmc-pb/v2/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE,PUT,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}

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
	builder.initCreds()
	return builder
}
func (b *Builder) initCreds() *Builder {
	if os.Getenv("PRODUCTION") == "1" {
		tlsCreds, err := loadClientTLSCredentials()
		if err != nil {
			panic(err)
		}
		b.opts = append(b.opts, grpc.WithTransportCredentials(tlsCreds))

	} else {
		b.opts = append(b.opts, grpc.WithInsecure())
	}
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
		if err := http.ListenAndServe("0.0.0.0:8080", cors(b.mux)); err != nil {
			panic(err)
		}
	}()

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	<-s
}
