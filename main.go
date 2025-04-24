package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alijabbar034/database"
	docs "github.com/alijabbar034/docs"
	"github.com/alijabbar034/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	docs.SwaggerInfo.BasePath = "/api"
}

// @title           Gin Ecommerce Web
// @version         1.0
// @description     An Ecommerce service API in Go using Gin framework.
// @contact.name   Ali Jabbar
// @contact.email  alijabbar0034@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      13.53.115.94
// @BasePath  /api

func main() {
	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Static("/resources", "./resources")

	evnLoadError := godotenv.Load()
	if evnLoadError != nil {
		log.Fatal("env Load Error")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001", "http://localhost:8000", "https://13.53.115.94/", "http://13.53.115.94/"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		MaxAge:       time.Hour * 12,
	}))

	// Initialize the database connection
	database.ConnectDB()

	// Register Swagger handler

	// Group your routes under /api
	api := app.Group("/api")
	api.GET("/", Welcome)
	routes.UserRouter(api)
	routes.ProductRouter(api)
	routes.ReviewRouter(api)
	routes.OrderRouter(api)

	port := os.Getenv("PORT")
	log.Fatal(app.Run(fmt.Sprintf(":%s", port)))
}

func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome",
	})

}
