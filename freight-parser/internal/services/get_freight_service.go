package services

import (
	"context"
	"errors"
	"fmc-newest/internal/logging"
	pb "fmc-newest/internal/proto"
	"fmc-newest/pkg/domain"
	"fmc-newest/pkg/freight"
	"fmt"
)

type converter struct {
}

func (c *converter) convertRequestToGetFreightStruct(request *pb.FreightRequest) domain.GetFreight {
	return domain.GetFreight{FromCity: request.FromCity.String(), ToCity: request.FromCity.String(), ContainerType: request.ContainerSize.String(), Limit: request.Limit}

}

func (c *converter) convertResponseToGrpcResponse(freights []domain.BaseFreight) *pb.FreightResponseList {
	var outputSlice []*pb.FreightResponse
	for _, value := range freights {
		oneGrpcResponse := pb.FreightResponse{
			FromCity: &pb.City{
				CityId:       int64(value.FromCity.CityId),
				CityName:     pb.CityEnum(pb.CityEnum_value[value.FromCity.CityName]),
				CityUnlocode: pb.Unlocode(pb.Unlocode_value[value.FromCity.CityUnlocode]),
			},
			ToCity: &pb.City{
				CityId:       int64(value.ToCity.CityId),
				CityName:     pb.CityEnum(pb.CityEnum_value[value.ToCity.CityName]),
				CityUnlocode: pb.Unlocode(pb.Unlocode_value[value.ToCity.CityUnlocode]),
			},
			ContainerType: &pb.Container{
				ContainerType:   pb.ContainerType(pb.ContainerType_value[value.ContainerType]),
				ContainerTypeId: int64(value.ContainerTypeId),
			},
			UsdPrice: int64(value.UsdPrice),
			Line: &pb.Line{
				LineId:    int64(value.LineId),
				Scac:      pb.ShippingLine(pb.ShippingLine_value[value.Line.LineName]),
				LineName:  value.LineName,
				LineImage: value.LineImage,
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
	return &pb.FreightResponseList{MultiResponse: outputSlice}
}

type GetFreightService struct {
	freightController freight.IController
	logger            logging.ILogger
	pb.UnimplementedFreightParserServer
	converter converter
}

func (s *GetFreightService) GetFreight(ctx context.Context, r *pb.FreightRequest) (*pb.FreightResponseList, error) {
	convertedRequest := s.converter.convertRequestToGetFreightStruct(r)
	ch := make(chan []domain.BaseFreight, 1)
	go func() {
		response, err := s.freightController.GetFreights(ctx, convertedRequest)
		if err != nil {
			s.logger.ExceptionLog(fmt.Sprintf(`GetFreights error: %s`, err.Error()))
		}
		ch <- response
	}()
	var freightResponseList *pb.FreightResponseList
	select {
	case <-ctx.Done():
		return freightResponseList, nil
	case result := <-ch:
		return s.converter.convertResponseToGrpcResponse(result), nil
	default:
		return freightResponseList, errors.New(fmt.Sprintf(`something went wrong,default case was selected`))
	}
}

func NewGetFreightService(freightController freight.IController, logger logging.ILogger) *GetFreightService {
	return &GetFreightService{
		freightController:                freightController,
		logger:                           logger,
		UnimplementedFreightParserServer: pb.UnimplementedFreightParserServer{},
	}
}

func NewConverter() *converter {
	return &converter{}
}
