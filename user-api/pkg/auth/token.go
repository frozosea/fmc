package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"UserId"`
}

type ITokenManager interface {
	GenerateToken(userId int, tokenExpiration int) (string, error)
	GenerateAccessRefreshTokens(userId int) (*Token, error)
	DecodeToken(tokenStr string) (int, error)
}

type TokenManager struct {
	SecretKey              string
	AccessTokenExpiration  int
	RefreshTokenExpiration int
}

func (t *TokenManager) GenerateToken(userId int, tokenExpiration int) (string, error) {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Duration(tokenExpiration) * time.Hour).Unix(),
		IssuedAt:  time.Now().Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{standardClaims, userId}).SignedString([]byte(t.SecretKey))
}

func (t *TokenManager) GenerateAccessRefreshTokens(userId int) (*Token, error) {
	var token Token
	accessToken, accessTokenGenerateErr := t.GenerateToken(userId, t.AccessTokenExpiration)
	if accessTokenGenerateErr != nil {
		return &token, accessTokenGenerateErr
	}
	refreshToken, refreshTokenGenerateErr := t.GenerateToken(userId, t.RefreshTokenExpiration)
	if refreshTokenGenerateErr != nil {
		return &token, refreshTokenGenerateErr
	}
	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return &token, nil

}
func (t *TokenManager) DecodeToken(tokenStr string) (int, error) {
	parsedToken, parsedTokenErr := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})
	if parsedTokenErr != nil {
		return 1, parsedTokenErr
	}
	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok {
		return 1, errors.New(`token claims are not valid`)
	}
	return claims.UserId, nil
}
