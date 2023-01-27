package util

import (
	"context"
	"errors"
	pb "github.com/frozosea/fmc-pb/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

type ITokenManager interface {
	GetUserIdFromCtx(ctx context.Context) (int, error)
	GetUserIdFromToken(ctx context.Context, token string) (int, error)
	GenerateGRPCAuthHeader(ctx context.Context, token string) (context.Context, grpc.CallOption)
	GetTokenFromHeaders(ctx context.Context) (string, error)
}

type TokenManager struct {
	client pb.AuthClient
}

func NewTokenManager(client pb.AuthClient) *TokenManager {
	return &TokenManager{client: client}
}

func (t *TokenManager) GetUserIdFromCtx(ctx context.Context) (int, error) {
	token, err := t.GetTokenFromHeaders(ctx)
	if err != nil {
		return -1, err
	}
	ctx, h := t.GenerateGRPCAuthHeader(ctx, token)
	response, err := t.client.GetUserIdByJwtToken(ctx, &emptypb.Empty{}, h)
	if err != nil {
		return -1, err
	}
	return int(response.GetUserId()), nil
}

func (t *TokenManager) GetUserIdFromToken(ctx context.Context, token string) (int, error) {
	ctx, h := t.GenerateGRPCAuthHeader(ctx, token)
	response, err := t.client.GetUserIdByJwtToken(ctx, &emptypb.Empty{}, h)
	if err != nil {
		return -1, err
	}
	return int(response.GetUserId()), nil
}

func (t *TokenManager) GenerateGRPCAuthHeader(ctx context.Context, token string) (context.Context, grpc.CallOption) {
	md := metadata.New(map[string]string{"ServerAuthorization": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, grpc.Header(&md)
}

func (t *TokenManager) GetTokenFromHeaders(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("not ok")
	}
	if md.Len() != 0 {
		token := md.Get("authorization")[0]
		if token == "" {
			return "", errors.New("not ok")
		}
		return strings.Split(token, " ")[0], nil
	}
	return "", errors.New("no authorization header")
}
