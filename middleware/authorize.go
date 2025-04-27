package middleware

import (
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Authorize  role
func Autorize(role string) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, exists := c.Get("user")
		if !exists {
			helper.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized access.ONLY admin can access it....")
			c.Abort()
			return
		}
		userObje, _ := user.(*models.User)
		if userObje.Role != role {
			helper.ErrorHandler(c, http.StatusUnauthorized, "Unauthorized access")
			c.Abort()
			return
		}
		c.Next()

	}
}
