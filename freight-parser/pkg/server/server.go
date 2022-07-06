package server

import (
	pb "fmc-newest/pkg/proto"
	"google.golang.org/grpc"
)

type server struct {
	freightService pb.FreightServiceServer
	cityService    pb.CityServiceServer
	contactService pb.ContactServiceServer
	lineService    pb.LineServiceServer
}

func (s *server) GetServer() *grpc.Server {
	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterFreightServiceServer(server, s.freightService)
	pb.RegisterCityServiceServer(server, s.cityService)
	pb.RegisterContactServiceServer(server, s.contactService)
	pb.RegisterLineServiceServer(server, s.lineService)
	return server
}
func NewServer(freightService pb.FreightServiceServer, cityService pb.CityServiceServer, contactService pb.ContactServiceServer, lineService pb.LineServiceServer) *server {
	return &server{freightService: freightService, cityService: cityService, contactService: contactService, lineService: lineService}
}
