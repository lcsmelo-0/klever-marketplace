package routes

import (
	"klever-marketplace/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoute(router *gin.Engine) {
	router.POST("/orders", controllers.CreateOrder())
	router.GET("/orders/:orderID", controllers.GetOrderByID())
	router.GET("/orders", controllers.GetAllOrders())
}
