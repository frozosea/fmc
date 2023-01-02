package utils

import (
	"context"
	"errors"
	"fmc-gateway/internal/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

type UserId int
type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"userId"`
}
type HttpUtils struct {
	cli *auth.Client
}

func NewHttpUtils(cli *auth.Client) *HttpUtils {
	return &HttpUtils{cli: cli}
}

func (h *HttpUtils) ValidateSchemaError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"success": false, "message": message})
}
func (h *HttpUtils) Validate(c *gin.Context, schema interface{}) error {
	if err := c.ShouldBindJSON(&schema); err != nil {
		return err
	}
	return nil
}
func (h *HttpUtils) GetUserIdByJwtToken(accessToken string) (UserId, error) {
	user, err := h.cli.GetUserIdByJwtToken(context.Background(), accessToken)
	if err != nil {
		return UserId(user), err
	}
	return UserId(user), nil
}
func (h *HttpUtils) DecodeToken(c *gin.Context) (UserId, error) {
	authHeader := c.GetHeader(`Authorization`)
	if authHeader == "" {
		return -1, errors.New("cannot decode token")
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 {
		return -1, errors.New("invalid token format")
	}
	userId, exc := h.GetUserIdByJwtToken(authParts[1])
	if exc != nil {
		return -1, exc
	}
	return userId, nil
}

func GenerateAuthHeader(ctx context.Context, token string) (context.Context, grpc.CallOption) {
	md := metadata.New(map[string]string{"Authorization": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, grpc.Header(&md)
}
