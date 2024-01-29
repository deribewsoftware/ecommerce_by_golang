package routes

import (
	"github.com/deribewsoftware/ecommerce_golang/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.Signup())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controllers.SearcProducts())--
	incomingRoutes.GET("/users/search", controllers.SearcProductsByQuery())

}
