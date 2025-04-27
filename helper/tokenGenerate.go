package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

// Generate jwt token and send via cookie
func SendToken(c *gin.Context, id string, message string) {
	envError := godotenv.Load()
	if envError != nil {
		ErrorHandler(c, http.StatusInternalServerError, "env get errr")
		return
	}
	SECRET_KEY := os.Getenv("SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id": id,
		"exp": time.Now().Add(time.Hour * 24 * 60).Unix(),
	})

	signedToken, eror := token.SignedString([]byte(SECRET_KEY))
	if eror != nil {
		ErrorHandler(c, http.StatusInternalServerError, "Token generation Eror")
		return
	}
	c.SetCookie("token", signedToken, int(time.Hour.Seconds()*24*60), "/", "localhost", false, true)
	c.JSON(http.StatusCreated, gin.H{
		"message": message,
		"token":   signedToken,
	})
}
