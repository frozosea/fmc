package auth

import (
	"context"
	"fmc-gateway/pkg/logging"
	"fmc-gateway/pkg/user-pb"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type converter struct{}

func (c *converter) convertUser(usr User) *__.LoginUserRequest {
	return &__.LoginUserRequest{
		Email:    usr.Email,
		Password: usr.Password,
	}
}
func (c *converter) convertLoginUser(usr *User) *__.LoginUserRequest {
	return &__.LoginUserRequest{
		Email:    usr.Email,
		Password: usr.Password,
	}
}
func (c *converter) convertRegisterUser(r *RegisterUser) *__.RegisterUserRequest {
	return &__.RegisterUserRequest{
		Email:    r.Email,
		Username: r.Username,
		Password: r.Password,
	}
}
func (c *converter) loginResponseConvert(response *__.LoginResponse) *LoginUserResponse {
	return &LoginUserResponse{
		Token:               response.GetTokens(),
		RefreshToken:        response.GetRefreshToken(),
		TokenType:           `Bearer`,
		TokenExpires:        time.Now().Add(time.Duration(response.GetTokenExpires()) * time.Hour).Unix(),
		RefreshTokenExpires: time.Now().Add(time.Duration(response.GetRefreshTokenExpires()) * time.Hour).Unix(),
	}
}

type AlreadyRegisterError struct {
}

func (a *AlreadyRegisterError) Error() string {
	return "user with this username already exists"
}

type InvalidUserError struct{}

func (i *InvalidUserError) Error() string {
	return "Cannot find user with these parameters"
}

type UnauthenticatedError struct{}

func (u *UnauthenticatedError) Error() string {
	return "Unauthenticated, cannot validate or decode token"
}

type IClient interface {
	Register(ctx context.Context, user *RegisterUser) error
	Login(ctx context.Context, user *User) (*LoginUserResponse, error)
	RefreshToken(ctx context.Context, refreshToken string) (*LoginUserResponse, error)
	CheckAccess(ctx context.Context, token string) (bool, error)
	GetUserIdByJwtToken(ctx context.Context, token string) (int64, error)
}
type Client struct {
	cli       __.AuthClient
	converter converter
	logger    logging.ILogger
}

func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{cli: __.NewAuthClient(conn), converter: converter{}, logger: logger}
}

func (c *Client) Register(ctx context.Context, user *RegisterUser) error {
	_, err := c.cli.RegisterUser(ctx, c.converter.convertRegisterUser(user))
	if err != nil {
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.AlreadyExists:
			return &AlreadyRegisterError{}
		default:
			return err
		}
	}
	return nil
}

func (c *Client) Login(ctx context.Context, user *User) (*LoginUserResponse, error) {
	r, err := c.cli.LoginUser(ctx, c.converter.convertLoginUser(user))
	if err != nil {
		statusCode := status.Convert(err).Code()
		switch statusCode {
		case codes.NotFound:
			return &LoginUserResponse{}, &InvalidUserError{}
		default:
			return &LoginUserResponse{}, err
		}
	}
	go c.logger.InfoLog(fmt.Sprintf(`user %s logged in successfully`, user.Email))
	return c.converter.loginResponseConvert(r), nil
}
func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (*LoginUserResponse, error) {
	r, err := c.cli.RefreshToken(ctx, &__.RefreshTokenRequest{RefreshToken: refreshToken})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`refresh failed: %s`, err.Error()))
		return &LoginUserResponse{}, err
	}
	return c.converter.loginResponseConvert(r), nil
}
func (c *Client) CheckAccess(ctx context.Context, token string) (bool, error) {
	success, err := c.cli.Auth(ctx, &__.AuthRequest{Tokens: token})
	if err != nil {
		return false, err
	}
	if !success.GetSuccess() {
		return false, nil
	}
	return success.GetSuccess(), nil
}
func (c *Client) GetUserIdByJwtToken(ctx context.Context, token string) (int64, error) {
	response, err := c.cli.GetUserIdByJwtToken(ctx, &__.GetUserIdByJwtTokenRequest{Token: token})
	statusCode := status.Convert(err).Code()
	switch statusCode {
	case codes.Unauthenticated:
		return -1, &UnauthenticatedError{}
	default:
		return response.GetUserId(), err
	}
}
