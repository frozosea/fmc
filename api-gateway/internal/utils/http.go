package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
	"strings"
)

type userId int
type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"UserId"`
}
type HttpUtils struct {
	secretKey string
}

func NewHttpUtils(secretKey string) *HttpUtils {
	return &HttpUtils{secretKey: secretKey}
}

func (h *HttpUtils) ValidateSchemaError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"success": false, "message": message})
}
func (h *HttpUtils) Validate(c *gin.Context, schema interface{}) error {
	if err := c.ShouldBindJSON(&schema); err != nil {
		if err := validator.Validate(schema); err != nil {
			return err
		}
		return err
	}
	return nil
}
func (h *HttpUtils) GetUserIdByJwtToken(accessToken string) (userId, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(h.secretKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New(`token claims are not valid`)
	}
	return userId(claims.UserId), nil
}
func (h *HttpUtils) DecodeToken(c *gin.Context) (userId, error) {
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
