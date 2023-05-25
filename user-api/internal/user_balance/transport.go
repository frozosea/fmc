package user_balance

import (
	"context"
	pb "github.com/frozosea/fmc-pb/v2/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/user_balance/credit"
	"user-api/pkg/util"
)

type Grpc struct {
	service       IService
	creditService credit.IService
	tokenManager  *util.TokenManager
	pb.UnimplementedBalanceServer
}

func NewGrpc(service IService, creditService credit.IService, tokenManager *util.TokenManager) *Grpc {
	return &Grpc{service: service, tokenManager: tokenManager, creditService: creditService, UnimplementedBalanceServer: pb.UnimplementedBalanceServer{}}
}

func (g *Grpc) SubOneDayTrackingPriceFromBalance(ctx context.Context, r *pb.SubBalanceServiceRequest) (*emptypb.Empty, error) {
	if err := g.service.SubOneDayTrackingPriceFromBalance(ctx, r.GetUserId(), r.GetNumber()); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
func (g *Grpc) GetTariff(ctx context.Context, _ *emptypb.Empty) (*pb.GetTariffResponse, error) {
	userId, err := g.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	t, err := g.service.GetCurrentTariff(ctx, int64(userId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetTariffResponse{
		OneDayTrackingPrice: float32(t.OneDayPrice),
		NumbersOnTrack:      t.NumbersOnTrackQuantity,
	}, nil
}
func (g *Grpc) GetBalance(ctx context.Context, _ *emptypb.Empty) (*pb.GetBalanceResponse, error) {
	userId, err := g.tokenManager.GetUserId(ctx)
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	b, err := g.service.GetBalance(ctx, int64(userId))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetBalanceResponse{Balance: float32(b)}, nil
}
func (g *Grpc) CheckAccessToPaidTracking(ctx context.Context, r *pb.CheckAccessToPaidTrackingRequest) (*pb.CheckAccessToPaidTrackingResponse, error) {
	hasAccess, err := g.creditService.CheckAccessToPaidTracking(ctx, r.GetUserId())
	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}
	return &pb.CheckAccessToPaidTrackingResponse{HasAccess: hasAccess}, nil
}
