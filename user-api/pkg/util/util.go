package util

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"user-api/internal/auth"
)

func GenerateGRPCAuthHeader(token string) grpc.CallOption {
	md := metadata.New(map[string]string{"Authorization": token})
	return grpc.Header(&md)
}

func GetTokenFromHeaders(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("not ok")
	}
	token := md.Get("Authorization")[0]
	if token == "" {
		return "", errors.New("not ok")
	}
	return token, nil
}

type TokenManager struct {
	manager auth.ITokenManager
}

func NewTokenManager(manager auth.ITokenManager) *TokenManager {
	return &TokenManager{manager: manager}
}

func (t *TokenManager) GetUserId(ctx context.Context) (int, error) {
	token, err := GetTokenFromHeaders(ctx)
	if err != nil {
		return -1, err
	}
	return t.manager.DecodeToken(token)
}
