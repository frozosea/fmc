package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Token struct {
	AccessToken            string        `json:"access_token"`
	RefreshToken           string        `json:"refresh_token"`
	TokenType              string        `json:"token_type"`
	AccessTokenExpiration  time.Duration `json:"access_token_expires"`
	RefreshTokenExpiration time.Duration `json:"refresh_token_expires"`
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}

type TokenClaimsWithOperationType struct {
	*TokenClaims
	OperationType string
}

type ITokenManager interface {
	GenerateToken(userId int, tokenExpiration time.Duration) (string, error)
	GenerateAccessRefreshTokens(userId int) (*Token, error)
	DecodeToken(tokenStr string) (int, error)
	GenerateResetPasswordToken(userId int, tokenExpiration time.Duration) (string, error)
	DecodeResetPasswordToken(token string) (int, string, error)
}

type TokenManager struct {
	SecretKey              string
	AccessTokenExpiration  time.Duration
	RefreshTokenExpiration time.Duration
}

func NewTokenManager(secretKey string, accessTokenExpiration time.Duration, refreshTokenExpiration time.Duration) *TokenManager {
	return &TokenManager{SecretKey: secretKey, AccessTokenExpiration: accessTokenExpiration, RefreshTokenExpiration: refreshTokenExpiration}
}

func (t *TokenManager) GenerateToken(userId int, tokenExpiration time.Duration) (string, error) {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenExpiration).Unix(),
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
	token.TokenType = "Bearer"
	token.AccessTokenExpiration = t.AccessTokenExpiration
	token.RefreshTokenExpiration = t.RefreshTokenExpiration
	return &token, nil

}
func (t *TokenManager) DecodeToken(tokenStr string) (int, error) {
	parsedToken, parsedTokenErr := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})
	if parsedTokenErr != nil {
		return -1, parsedTokenErr
	}
	claims, ok := parsedToken.Claims.(*TokenClaims)
	if !ok {
		return -1, errors.New(`token claims are not valid`)
	}
	return claims.UserId, nil
}
func (t *TokenManager) GenerateResetPasswordToken(userId int, tokenExpiration time.Duration) (string, error) {
	standardClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenExpiration).Unix(),
		IssuedAt:  time.Now().Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsWithOperationType{TokenClaims: &TokenClaims{StandardClaims: standardClaims, UserId: userId}, OperationType: "reset_password"}).SignedString([]byte(t.SecretKey))
}
func (t *TokenManager) DecodeResetPasswordToken(token string) (int, string, error) {
	parsedToken, parsedTokenErr := jwt.ParseWithClaims(token, &TokenClaimsWithOperationType{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})
	if parsedTokenErr != nil {
		return -1, "", parsedTokenErr
	}
	claims, ok := parsedToken.Claims.(*TokenClaimsWithOperationType)
	if !ok {
		return -1, "", errors.New(`token claims are not valid`)
	}
	return claims.UserId, claims.OperationType, nil
}
