package routes

import (
	"github.com/menacedjava/controllers"
	"github.com/menacedjava/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	routerGroup := router.Group("/user")
	routerGroup.GET("/", middleware.Authenticate, middleware.Autorize("admin"),controllers.GetAllUsers)
	routerGroup.GET("/:id", middleware.Authenticate, middleware.Autorize("admin"),controllers.GetUser)
	routerGroup.POST("/register", controllers.RegisterUser)
	routerGroup.POST("/login", controllers.LoginUser)
	routerGroup.GET("/me", middleware.Authenticate, controllers.GetProfile)
	routerGroup.PUT("/me", middleware.Authenticate, controllers.UpdateUser)
	routerGroup.GET("/logout", middleware.Authenticate, controllers.LogoutUser)
}

