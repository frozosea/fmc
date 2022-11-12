package feedback

import (
	"github.com/gin-gonic/gin"
	"os"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) CheckAdminAccess(c *gin.Context) {
	adminPassword := os.Getenv("ADMIN_ACCESS_KEY")
	if adminPassword == "" {
		c.JSON(401, gin.H{"success": false, "error": "missing auth!"})
		return
	}
	k := c.Request.Header.Get("Authorization")
	if k == "" {
		c.JSON(401, gin.H{"success": false, "error": "missing auth!"})
		return
	}
	if k != adminPassword {
		c.JSON(401, gin.H{"success": false, "error": "missing auth!"})
		return
	}
}
