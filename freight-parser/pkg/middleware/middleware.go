package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Middleware struct {
}

func New() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CheckAuth(c *gin.Context) {
	authHeader := c.GetHeader(`Authorization`)
	if authHeader == "" {
		c.AbortWithStatus(401)
		return
	}
	//TODO put in docker and markdown and .env
	accessPassword := os.Getenv("ACCESS_PASSWORD")
	if accessPassword == "" {
		c.AbortWithStatus(401)
		return
	}
	if accessPassword != authHeader {
		c.AbortWithStatus(401)
		return
	}
}
