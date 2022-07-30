package freight

import (
	"context"
	"fmc-newest/internal/logging"
	pb "fmc-newest/pkg/proto"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type adapter struct {
}

func (a *adapter) convertRequestToGetFreightStruct(request *pb.GetFreightRequest) GetFreight {
	return GetFreight{FromCityId: request.GetFromCityId(), ToCityId: request.GetToCityId(), ContainerTypeId: request.GetContainerTypeId(), Limit: request.GetLimit()}

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
func (a *adapter) convertAddFreight(r *pb.AddFreightRequest) AddFreight {
	return AddFreight{
		FromCityId:      r.GetFromCityId(),
		ToCityId:        r.GetToCityId(),
		ContainerTypeId: int(r.GetContainerTypeId()),
		UsdPrice:        int(r.GetUsdPrice()),
		LineId:          r.GetLineId(),
		FromDate:        time.Unix(r.GetFromDate().GetSeconds(), 0),
		ExpiresDate:     time.Unix(r.GetExpiryDate().GetSeconds(), 0),
		ContactId:       int(r.GetContactId()),
	}
}

type Service struct {
	controller *controller
	logger     logging.ILogger
	pb.UnimplementedFreightServiceServer
	converter adapter
}

func (s *Service) GetFreight(ctx context.Context, r *pb.GetFreightRequest) (*pb.GetFreightsResponseList, error) {
	response, err := s.controller.GetFreights(ctx, s.converter.convertRequestToGetFreightStruct(r))
	if err != nil {
		s.logger.ExceptionLog(fmt.Sprintf(`GetFreights error: %s`, err.Error()))
		return &pb.GetFreightsResponseList{}, status.Error(codes.Internal, err.Error())
	}
	return s.converter.convertResponseToGrpcResponse(response), nil
}
func (s *Service) AddFreight(ctx context.Context, r *pb.AddFreightRequest) (*emptypb.Empty, error) {
	if err := s.controller.AddFreight(ctx, s.converter.convertAddFreight(r)); err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func NewGetFreightService(freightController *controller, logger logging.ILogger) *Service {
	return &Service{
		controller:                        freightController,
		logger:                            logger,
		UnimplementedFreightServiceServer: pb.UnimplementedFreightServiceServer{},
	}
}

func NewConverter() *adapter {
	return &adapter{}
}
