package server

import "google.golang.org/grpc"

func GetServer() *grpc.Server {
	server := grpc.NewServer(grpc.EmptyServerOption{})
	return server
}
