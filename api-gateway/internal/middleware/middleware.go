package middleware

import (
	"fmc-with-git/internal/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type Middleware struct {
	*utils.HttpUtils
}

func NewMiddleware(httpUtils *utils.HttpUtils) *Middleware {
	return &Middleware{HttpUtils: httpUtils}
}

func (m *Middleware) CheckAccessMiddleware(c *gin.Context) {
	authHeader := c.GetHeader(`Authorization`)
	if authHeader == "" {
		c.AbortWithStatus(401)
		return
	}
	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 {
		c.AbortWithStatus(401)
		return
	}
	_, exc := m.GetUserIdByJwtToken(authParts[1])
	if exc != nil {
		fmt.Println(exc.Error())
		c.AbortWithStatus(401)
	}
}
