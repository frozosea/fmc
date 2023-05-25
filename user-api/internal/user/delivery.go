package user

import (
	"context"
	pb "github.com/frozosea/fmc-pb/v2/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/domain"
	"user-api/pkg/util"
)

type converter struct{}

func (c *converter) convertContainerOrBillToGrpc(r []*domain.Container) []*pb.ContainerResponse {
	var arr []*pb.ContainerResponse
	for _, v := range r {
		pbResp := &pb.ContainerResponse{
			Number:      v.Number,
			IsOnTrack:   v.IsOnTrack,
			IsContainer: v.IsContainer,
		}
		if v.IsOnTrack {
			pbResp.ScheduleTrackingObject = &pb.ScheduleTrackingObject{
				Time:    v.ScheduleTrackingInfo.Time,
				Emails:  v.ScheduleTrackingInfo.Emails,
				Subject: v.ScheduleTrackingInfo.Subject,
			}
		}
		arr = append(arr, pbResp)
	}
	return arr
}
func (c *converter) addContainerOrBillConvert(r *pb.AddContainerToAccountRequest) []string {
	var containers []string
	for _, v := range r.GetContainer() {
		containers = append(containers, v.GetNumber())
	}
	return containers
}

func (c *converter) getInfoAboutUserConvert(user *domain.UserWithId) *pb.GetInfoAboutUserResponse {
	return &pb.GetInfoAboutUserResponse{
		Id:                     int64(user.Id),
		Email:                  user.Email,
		Username:               user.Username,
		CompanyFullName:        user.CompanyData.CompanyFullName,
		CompanyAbbreviatedName: user.CompanyData.CompanyAbbreviatedName,
		Inn:                    user.CompanyData.INN,
		Ogrn:                   user.CompanyData.OGRN,
		LegalAddress:           user.CompanyData.LegalAddress,
		PostAddress:            user.CompanyData.PostAddress,
		WorkEmail:              user.CompanyData.WorkEmail,
		Tariff: &pb.Tariff{
			NumbersOnTracking:   user.Tariff.NumbersOnTrackQuantity,
			OneDayTrackingPrice: float32(user.Tariff.OneDayPrice),
		},
		Numbers: &pb.GetAllContainersResponse{
			BillNumbers: c.convertContainerOrBillToGrpc(user.Numbers.BillNumbers),
			Containers:  c.convertContainerOrBillToGrpc(user.Numbers.Containers),
		},
	}
}

func (c *converter) convertUpdateCompanyDataRequest(r *pb.UpdateCompanyDataRequest) *domain.CompanyData {
	return &domain.CompanyData{
		CompanyFullName:        r.GetCompanyFullName(),
		CompanyAbbreviatedName: r.GetCompanyAbbreviatedName(),
		INN:                    r.GetInn(),
		OGRN:                   r.GetOgrn(),
		LegalAddress:           r.GetLegalAddress(),
		PostAddress:            r.GetPostAddress(),
		WorkEmail:              r.GetWorkEmail(),
	}
}

type Grpc struct {
	service      *Service
	converter    converter
	tokenManager *util.TokenManager
	pb.UnimplementedUserServer
}

func NewGrpc(controller *Service, manager *util.TokenManager) *Grpc {
	return &Grpc{
		service:                 controller,
		converter:               converter{},
		tokenManager:            manager,
		UnimplementedUserServer: pb.UnimplementedUserServer{},
	}
}

func (s *Grpc) AddContainerToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.AddContainerToAccount(ctx, userId, s.converter.addContainerOrBillConvert(r)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (s *Grpc) AddBillNumberToAccount(ctx context.Context, r *pb.AddContainerToAccountRequest) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.AddBillNumberToAccount(ctx, userId, s.converter.addContainerOrBillConvert(r)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (s *Grpc) DeleteContainersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.DeleteContainersFromAccount(ctx, userId, r.GetNumbers()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil

}
func (s *Grpc) DeleteBillNumbersFromAccount(ctx context.Context, r *pb.DeleteContainersFromAccountRequest) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.DeleteBillNumbersFromAccount(ctx, userId, r.GetNumbers()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil

}
func (s *Grpc) GetAll(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllContainersResponse, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	res, err := s.service.GetAllContainers(ctx, userId)
	if err != nil {
		return &pb.GetAllContainersResponse{}, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetAllContainersResponse{
		BillNumbers: s.converter.convertContainerOrBillToGrpc(res.BillNumbers),
		Containers:  s.converter.convertContainerOrBillToGrpc(res.Containers),
	}, nil
}
func (s *Grpc) GetInfoAboutUser(ctx context.Context, _ *emptypb.Empty) (*pb.GetInfoAboutUserResponse, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}
	data, err := s.service.GetInfoAboutUser(ctx, userId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return s.converter.getInfoAboutUserConvert(data), nil
}
func (s *Grpc) UpdateCompanyData(ctx context.Context, r *pb.UpdateCompanyDataRequest) (*emptypb.Empty, error) {
	userId, err := s.tokenManager.GetUserId(ctx)
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Unauthenticated, err.Error())
	}
	if err := s.service.UpdateCompanyData(ctx, userId, s.converter.convertUpdateCompanyDataRequest(r)); err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
