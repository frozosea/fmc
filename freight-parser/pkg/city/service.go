package city

import (
	"context"
	"fmc-newest/pkg/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

type cityConverter struct {
}

func (c *cityConverter) addCityGrpcRequestConvertToAddCityStruct(addCityRequest *___.AddCityRequest) *BaseCity {
	return &BaseCity{FullName: addCityRequest.CityFullName, Unlocode: addCityRequest.Unlocode}
}
func (c *cityConverter) convertResponseToGrpcResponse(cities []*City) *___.GetAllCitiesResponse {
	var outputCitiesArray []*___.City
	for _, city := range cities {
		oneGrpcCity := ___.City{CityId: int64(city.Id), CityName: city.FullName, CityUnlocode: city.Unlocode}
		outputCitiesArray = append(outputCitiesArray, &oneGrpcCity)
	}
	return &___.GetAllCitiesResponse{Cities: outputCitiesArray}
}

type Service struct {
	controller IController
	___.UnimplementedCityServiceServer
	converter *cityConverter
}

func (s *Service) AddCity(ctx context.Context, addCityRequest *___.AddCityRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.controller.AddCity(ctx, *s.converter.addCityGrpcRequestConvertToAddCityStruct(addCityRequest))
}
func (s *Service) GetAllCities(ctx context.Context, _ *emptypb.Empty) (*___.GetAllCitiesResponse, error) {
	result, err := s.controller.GetAll(ctx)
	return s.converter.convertResponseToGrpcResponse(result), err
}
