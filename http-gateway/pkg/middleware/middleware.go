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

func (m *Middleware) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Cache-Control, Content-Type, DNT, If-Modified-Since, Keep-Alive, Origin, User-Agent, X-Requested-With, X-Real-Ip, Access-Control-Allow-Origin")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.Writer.Header().Set("Access-Control-Max-Age", "1728000")
		c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
		c.Writer.Header().Set("Content-Length", "0")
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
