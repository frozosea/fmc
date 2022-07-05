package auth

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"user-api/internal/logging"
	"user-api/pkg/domain"
	pb "user-api/pkg/proto"
)

type converter struct{}

func (c *converter) registerUserConvert(r *pb.RegisterUserRequest) domain.User {
	return domain.User{
		Username: r.GetUsername(),
		Password: r.GetPassword(),
	}
}
func (c *converter) loginUserConvert(r *pb.LoginUserRequest) domain.User {
	return domain.User{
		Username: r.GetUsername(),
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
	controller *Controller
	converter  converter
	logger     logging.ILogger
	pb.UnimplementedAuthServer
}

func NewService(controller *Controller, logger logging.ILogger) *Service {
	return &Service{controller: controller, converter: converter{}, logger: logger, UnimplementedAuthServer: pb.UnimplementedAuthServer{}}
}

func (s *Service) RegisterUser(ctx context.Context, r *pb.RegisterUserRequest) (*emptypb.Empty, error) {
	if err := s.controller.RegisterUser(ctx, s.converter.registerUserConvert(r)); err != nil {
		//go s.logger.ExceptionLog(fmt.Sprintf(`register user-pb with username "%s" error: %s`, r.GetUsername(), err.Error()))
		switch err.(type) {
		case *AlreadyRegisterError:
			return &emptypb.Empty{}, nil
		default:
			return &emptypb.Empty{}, err
		}
	}
	//go s.logger.InfoLog(fmt.Sprintf(`user-pb with username "%s" was registered`, r.GetUsername()))
	return &emptypb.Empty{}, nil
}

func (s *Service) LoginUser(ctx context.Context, r *pb.LoginUserRequest) (*pb.LoginResponse, error) {
	resp, err := s.controller.Login(ctx, s.converter.loginUserConvert(r))
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	//go s.logger.InfoLog(fmt.Sprintf(`login user-pb with username "%s" was login`, r.GetUsername()))
	return s.converter.loginResponseConvert(resp), nil
}
func (s *Service) RefreshToken(ctx context.Context, r *pb.RefreshTokenRequest) (*pb.LoginResponse, error) {
	resp, err := s.controller.RefreshToken(r.GetRefreshToken())
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	return s.converter.loginResponseConvert(resp), nil
}
func (s *Service) Auth(ctx context.Context, r *pb.AuthRequest) (*pb.AuthResponse, error) {
	success, err := s.controller.CheckAccess(ctx, r.GetTokens())
	if err != nil || !success {
		return &pb.AuthResponse{Success: false}, err
	}
	return &pb.AuthResponse{Success: true}, nil

}
