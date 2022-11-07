package middleware

import (
	"fmc-gateway/internal/auth"
	"fmc-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type Middleware struct {
	cli *auth.Client
	*utils.HttpUtils
}

func NewMiddleware(httpUtils *utils.HttpUtils, cli *auth.Client) *Middleware {
	return &Middleware{HttpUtils: httpUtils, cli: cli}
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
	hasAccess, err := m.cli.CheckAccess(c.Request.Context(), authParts[1])
	if err != nil || !hasAccess {
		c.AbortWithStatus(401)
		return
	}
}
