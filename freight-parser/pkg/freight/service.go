package freight

import (
	"context"
	"fmc-newest/internal/logging"
	pb "fmc-newest/pkg/proto"
	"fmt"
)

type adapter struct {
}

func (a *adapter) convertRequestToGetFreightStruct(request *pb.GetFreightRequest) GetFreight {
	return GetFreight{FromCityId: request.FromCityId, ToCityId: request.ToCityId, ContainerType: request.ContainerType.String(), Limit: request.Limit}

}

func (a *adapter) convertResponseToGrpcResponse(freights []BaseFreight) *pb.GetFreightsResponseList {
	var outputSlice []*pb.GetFreightResponse
	for _, value := range freights {
		oneGrpcResponse := pb.GetFreightResponse{
			FromCity: &pb.City{
				CityId:       int64(value.FromCity.Id),
				CityName:     value.FromCity.FullName,
				CityUnlocode: value.FromCity.Unlocode,
			},
			ToCity: &pb.City{
				CityId:       int64(value.ToCity.Id),
				CityName:     value.ToCity.FullName,
				CityUnlocode: value.ToCity.Unlocode,
			},
			ContainerType: &pb.Container{
				ContainerType:   pb.ContainerType(pb.ContainerType_value[value.Type]),
				ContainerTypeId: int64(value.Id),
			},
			UsdPrice: int64(value.UsdPrice),
			Line: &pb.Line{
				LineId:    int64(value.Id),
				Scac:      value.Line.Scac,
				LineName:  value.FullName,
				LineImage: value.ImageUrl,
			},
			Contact: &pb.Contact{
				Url:         value.Contact.Url,
				PhoneNumber: value.Contact.PhoneNumber,
				AgentName:   value.Contact.AgentName,
				Email:       value.Contact.Email,
			},
		}
		outputSlice = append(outputSlice, &oneGrpcResponse)
	}
	return &pb.GetFreightsResponseList{MultiResponse: outputSlice}
}

type GetFreightService struct {
	controller IController
	logger     logging.ILogger
	pb.UnimplementedFreightServiceServer
	converter adapter
}

func (s *GetFreightService) GetFreight(ctx context.Context, r *pb.GetFreightRequest) (*pb.GetFreightsResponseList, error) {
	convertedRequest := s.converter.convertRequestToGetFreightStruct(r)
	response, err := s.controller.GetFreights(ctx, convertedRequest)
	if err != nil {
		s.logger.ExceptionLog(fmt.Sprintf(`GetFreights error: %s`, err.Error()))
		return s.converter.convertResponseToGrpcResponse(response), err
	}
	return s.converter.convertResponseToGrpcResponse(response), nil
}

func NewGetFreightService(freightController IController, logger logging.ILogger) *GetFreightService {
	return &GetFreightService{
		controller:                        freightController,
		logger:                            logger,
		UnimplementedFreightServiceServer: pb.UnimplementedFreightServiceServer{},
	}
}

func NewConverter() *adapter {
	return &adapter{}
}
