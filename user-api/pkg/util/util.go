package util

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"strings"
)

func GetTokenFromHeaders(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", errors.New("not ok")
	}
	if md.Len() > 0 {
		tokenAr := md.Get("authorization")
		if len(tokenAr) > 0 {
			token := tokenAr[0]
			if token == "" {
				return "", errors.New("not ok")
			}
			return strings.Split(token, " ")[1], nil
		}
	}
	return "", errors.New("len")
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
