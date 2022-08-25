package routes

import (
	"klever-marketplace/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/influencer", controllers.CreateInfluencer())
}
