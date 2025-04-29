package routes

import (
	"github.com/menacedjava/controllers"
	"github.com/menacedjava/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRouter(c *gin.RouterGroup) {
	routes := c.Group("/order")
	routes.POST("/shipping", controllers.CreateShiipingAddress)
	routes.POST("/", controllers.CreateOrder)
	routes.GET("/", middleware.Authenticate, middleware.Autorize("admin"), controllers.GetAllOrders)
	routes.GET("/:id", middleware.Authenticate, middleware.Autorize("admin"), controllers.GetOrder)
	routes.PUT("/:id", middleware.Authenticate, middleware.Autorize("admin"), controllers.UpdateOrde)

}

