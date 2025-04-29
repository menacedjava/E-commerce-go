package routes

import (
	"github.com/alijabbar034/controllers"
	"github.com/alijabbar034/middleware"
	"github.com/gin-gonic/gin"
)

func ProductRouter(r *gin.RouterGroup) {
	router := r.Group("/product")

	router.GET("/", controllers.GetAllProduct)
	router.POST("/", middleware.Authenticate, middleware.Autorize("admin"), controllers.CreateProduct)
	router.GET("/:id", controllers.GetProductById)
	router.PUT("/:id", middleware.Authenticate, middleware.Autorize("admin"), controllers.UpdateProduct)
	router.DELETE("/:id", middleware.Authenticate, middleware.Autorize("admin"), controllers.DeleteProduct)
}
