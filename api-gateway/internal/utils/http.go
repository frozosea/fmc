package utils

import "github.com/gin-gonic/gin"

type HttpUtils struct {
}

func (s *HttpUtils) ValidateSchemaError(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, gin.H{"success": false, "message": message})
}
