package city

import (
	"context"
	pb "fmc-newest/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cityConverter struct {
}

func (c *cityConverter) addCityGrpcRequestConvertToAddCityStruct(addCityRequest *pb.AddCityRequest) *BaseCity {
	return &BaseCity{FullName: addCityRequest.CityFullName, Unlocode: addCityRequest.Unlocode}
}
func (c *cityConverter) convertResponseToGrpcResponse(cities []*City) *pb.GetAllCitiesResponse {
	var outputCitiesArray []*pb.City
	for _, city := range cities {
		oneGrpcCity := pb.City{CityId: int64(city.Id), CityName: city.FullName, CityUnlocode: city.Unlocode}
		outputCitiesArray = append(outputCitiesArray, &oneGrpcCity)
	}
	return &pb.GetAllCitiesResponse{Cities: outputCitiesArray}
}

type Service struct {
	controller IController
	pb.UnimplementedCityServiceServer
	converter *cityConverter
}

func (s *Service) AddCity(ctx context.Context, addCityRequest *pb.AddCityRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddCity(ctx, *s.converter.addCityGrpcRequestConvertToAddCityStruct(addCityRequest))
}
func (s *Service) GetAllCities(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllCitiesResponse, error) {
	result, err := s.controller.GetAll(ctx)
	return s.converter.convertResponseToGrpcResponse(result), err
}
