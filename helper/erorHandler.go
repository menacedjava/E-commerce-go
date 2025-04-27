package helper

import "github.com/gin-gonic/gin"

// Eror Handler  function
func ErrorHandler(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{
		"message": message,
	})
	return
}

