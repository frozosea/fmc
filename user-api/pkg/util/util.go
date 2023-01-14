package util

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
)

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
	decode func(tokenStr string) (int, error)
}

func NewTokenManager(decode func(tokenStr string) (int, error)) *TokenManager {
	return &TokenManager{decode: decode}
}

func (t *TokenManager) GetUserId(ctx context.Context) (int, error) {
	token, err := GetTokenFromHeaders(ctx)
	if err != nil {
		return -1, err
	}
	return t.decode(token)
}
