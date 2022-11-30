package auth

import (
	"context"
	"fmt"
	pb "github.com/frozosea/fmc-pb/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/domain"
	"user-api/pkg/logging"
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

type Grpc struct {
	controller *Service
	converter  converter
	logger     logging.ILogger
	pb.UnimplementedAuthServer
}

func NewGrpc(controller *Service, logger logging.ILogger) *Grpc {
	return &Grpc{controller: controller, converter: converter{}, logger: logger, UnimplementedAuthServer: pb.UnimplementedAuthServer{}}
}

func (s *Grpc) RegisterUser(ctx context.Context, r *pb.RegisterUserRequest) (*emptypb.Empty, error) {
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

func (s *Grpc) LoginUser(ctx context.Context, r *pb.LoginUserRequest) (*pb.LoginResponse, error) {
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
func (s *Grpc) RefreshToken(ctx context.Context, r *pb.RefreshTokenRequest) (*pb.LoginResponse, error) {
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
func (s *Grpc) Auth(ctx context.Context, r *pb.AuthRequest) (*pb.AuthResponse, error) {
	success, err := s.controller.CheckAccess(ctx, r.GetTokens())
	if err != nil || !success {
		return &pb.AuthResponse{Success: false}, status.Error(codes.Internal, err.Error())
	}
	return &pb.AuthResponse{Success: true}, nil

}
func (s *Grpc) GetUserIdByJwtToken(ctx context.Context, r *pb.GetUserIdByJwtTokenRequest) (*pb.GetUserIdByJwtTokenResponse, error) {
	userId, err := s.controller.GetUserIdByJwtToken(r.GetToken())
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "cannot decode jwt token ")
	}
	return &pb.GetUserIdByJwtTokenResponse{UserId: int64(userId)}, nil
}
func (s *Grpc) SendRecoveryEmail(ctx context.Context, r *pb.SendRecoveryEmailRequest) (*emptypb.Empty, error) {
	err := s.controller.SendRecoveryUserEmail(ctx, r.GetEmail())
	if err != nil {
		switch err.(type) {
		case *InvalidUserError:
			return &emptypb.Empty{}, status.Error(codes.NotFound, err.Error())
		default:
			return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
		}

	}
	return &emptypb.Empty{}, nil

}
func (s *Grpc) RecoveryUser(ctx context.Context, r *pb.RecoveryUserRequest) (*emptypb.Empty, error) {
	err := s.controller.RecoveryUser(ctx, r.GetToken(), r.GetPassword())
	if err != nil {
		return &emptypb.Empty{}, status.Error(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil

}
