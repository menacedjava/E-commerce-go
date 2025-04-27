package middleware

import (
	"fmt"
	"github.com/menacedjava/helper"
	"github.com/menacedjava/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

// Check if user has valid token or not
func Authenticate(c *gin.Context) {

	tokenString, err := c.Cookie("token")

	if err != nil {
		helper.ErrorHandler(c, http.StatusUnauthorized, "unauthorized access,.")
		c.Abort()
		return
	}
	envError := godotenv.Load()
	if envError != nil {
		helper.ErrorHandler(c, http.StatusInternalServerError, "env get errr")
		return
	}
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token, er := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if er != nil || !token.Valid {
		helper.ErrorHandler(c, http.StatusUnauthorized, "unauthorized access.No token valid")
		c.Abort()
		return
	}
	claims := token.Claims.(*jwt.MapClaims)
	userID, exists := (*claims)["_id"].(string)
	if !exists {
		helper.ErrorHandler(c, http.StatusUnauthorized, "unauthorized access.Id")
		c.Abort()
		return
	}
	fmt.Println("uuuuuuuuuuuuuu", userID)
	user, erorr := models.GetById(userID)
	fmt.Println("userId", user)
	if erorr != nil {
		helper.ErrorHandler(c, http.StatusUnauthorized, "unauthorized access,............")
		c.Abort()
		return
	}

	c.Set("user", user)
	c.Next()
}

