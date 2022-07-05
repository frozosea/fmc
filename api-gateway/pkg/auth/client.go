package auth

import (
	"context"
	"errors"
	"fmc-with-git/internal/logging"
	pb "fmc-with-git/internal/user-pb"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type converter struct{}

func (c *converter) convertUser(usr User) *pb.RegisterUserRequest {
	return &pb.RegisterUserRequest{
		Username: usr.Username,
		Password: usr.Password,
	}
}
func (c *converter) convertLoginUser(usr User) *pb.LoginUserRequest {
	return &pb.LoginUserRequest{
		Username: usr.Username,
		Password: usr.Password,
	}
}
func (c *converter) loginResponseConvert(response *pb.LoginResponse) *LoginUserResponse {
	return &LoginUserResponse{
		Token:               response.GetTokens(),
		RefreshToken:        response.GetRefreshToken(),
		TokenType:           `Bearer`,
		TokenExpires:        time.Now().Add(time.Duration(response.GetTokenExpires()) * time.Hour).Unix(),
		RefreshTokenExpires: time.Now().Add(time.Duration(response.GetRefreshTokenExpires()) * time.Hour).Unix(),
	}
}

type Client struct {
	cli       pb.AuthClient
	converter converter
	logger    logging.ILogger
}

func NewClient(conn *grpc.ClientConn, logger logging.ILogger) *Client {
	return &Client{cli: pb.NewAuthClient(conn), converter: converter{}, logger: logger}
}

func (c *Client) Register(ctx context.Context, user User) error {
	fmt.Println(user)
	_, err := c.cli.RegisterUser(ctx, c.converter.convertUser(user))
	if err != nil {
		//go c.logger.ExceptionLog(fmt.Sprintf(`register user with username %s failed: %s`, user.Username, err.Error()))
		return errors.New("cannot register user with username " + user.Username)
	}
	return nil
}

func (c *Client) Login(ctx context.Context, user User) (*LoginUserResponse, error) {
	r, err := c.cli.LoginUser(ctx, c.converter.convertLoginUser(user))
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`login user with username %s failed: %s`, user.Username, err.Error()))
		return &LoginUserResponse{}, err
	}
	go c.logger.InfoLog(fmt.Sprintf(`user %s logged in successfully`, user.Username))
	return c.converter.loginResponseConvert(r), nil
}
func (c *Client) RefreshToken(ctx context.Context, refreshToken string) (*LoginUserResponse, error) {
	r, err := c.cli.RefreshToken(ctx, &pb.RefreshTokenRequest{RefreshToken: refreshToken})
	if err != nil {
		go c.logger.ExceptionLog(fmt.Sprintf(`refresh failed: %s`, err.Error()))
		return &LoginUserResponse{}, err
	}
	return c.converter.loginResponseConvert(r), nil
}
func (c *Client) CheckAccess(ctx context.Context, token string) (bool, error) {
	success, err := c.cli.Auth(ctx, &pb.AuthRequest{Tokens: token})
	if err != nil {
		return false, err
	}
	if !success.GetSuccess() {
		return false, nil
	}
	return success.GetSuccess(), nil
}
