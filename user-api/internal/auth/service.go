package auth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/domain"
	"user-api/pkg/logging"
	pb "user-api/pkg/proto"
)

type converter struct{}

func (c *converter) registerUserConvert(r *pb.RegisterUserRequest) *domain.RegisterUser {
	return &domain.RegisterUser{
		Email:    r.GetEmail(),
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}
}
func (c *converter) loginUserConvert(r *pb.LoginUserRequest) *domain.User {
	return &domain.User{
		Email:    r.GetEmail(),
		Password: r.GetPassword(),
	}
}
func (c *converter) loginResponseConvert(t *Token) *pb.LoginResponse {
	return &pb.LoginResponse{
		Tokens:              t.AccessToken,
		RefreshToken:        t.RefreshToken,
		TokenExpires:        int64(t.AccessTokenExpiration),
		RefreshTokenExpires: int64(t.RefreshTokenExpiration),
	}
}

type Service struct {
	controller *Provider
	converter  converter
	logger     logging.ILogger
	pb.UnimplementedAuthServer
}

func NewService(controller *Provider, logger logging.ILogger) *Service {
	return &Service{controller: controller, converter: converter{}, logger: logger, UnimplementedAuthServer: pb.UnimplementedAuthServer{}}
}

func (s *Service) RegisterUser(ctx context.Context, r *pb.RegisterUserRequest) (*emptypb.Empty, error) {
	if err := s.controller.RegisterUser(ctx, s.converter.registerUserConvert(r)); err != nil {
		switch err.(type) {
		case *AlreadyRegisterError:
			return &emptypb.Empty{}, status.Error(codes.AlreadyExists, "user with these parameters already exists")
		default:
			return &emptypb.Empty{}, err
		}
	}
	go s.logger.InfoLog(fmt.Sprintf(`user-pb with username "%s" was registered`, r.GetUsername()))
	return &emptypb.Empty{}, nil
}

func (s *Service) LoginUser(ctx context.Context, r *pb.LoginUserRequest) (*pb.LoginResponse, error) {
	resp, err := s.controller.Login(ctx, s.converter.loginUserConvert(r))
	if err != nil {
		switch err.(type) {
		case *InvalidUserError:
			return &pb.LoginResponse{}, status.Error(codes.NotFound, "cannot login user with these parameters")
		default:
			return &pb.LoginResponse{}, status.Error(codes.Internal, err.Error())
		}
	}
	return s.converter.loginResponseConvert(resp), nil
}
func (s *Service) RefreshToken(ctx context.Context, r *pb.RefreshTokenRequest) (*pb.LoginResponse, error) {
	ch := make(chan *pb.LoginResponse, 1)
	errCh := make(chan error, 1)
	go func() {
		resp, err := s.controller.RefreshToken(r.GetRefreshToken())
		if err != nil {
			errCh <- status.Error(codes.Internal, err.Error())
			ch <- &pb.LoginResponse{}
		}
		errCh <- nil
		ch <- s.converter.loginResponseConvert(resp)
	}()
	select {
	case <-ctx.Done():
		return &pb.LoginResponse{}, ctx.Err()
	case result := <-ch:
		return result, <-errCh
	}
}
func (s *Service) Auth(ctx context.Context, r *pb.AuthRequest) (*pb.AuthResponse, error) {
	success, err := s.controller.CheckAccess(ctx, r.GetTokens())
	if err != nil || !success {
		return &pb.AuthResponse{Success: false}, err
	}
	return &pb.AuthResponse{Success: true}, nil

}
func (s *Service) GetUserIdByJwtToken(ctx context.Context, r *pb.GetUserIdByJwtTokenRequest) (*pb.GetUserIdByJwtTokenResponse, error) {
	userId, err := s.controller.GetUserIdByJwtToken(ctx, r.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "cannot decode jwt token ")
	}
	return &pb.GetUserIdByJwtTokenResponse{UserId: int64(userId)}, nil
}
