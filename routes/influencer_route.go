package routes

import (
	"klever-marketplace/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/influencers", controllers.CreateInfluencer())
	router.GET("/influencers/:influencerId", controllers.GetInfluencerByID())
	router.GET("/influencers", controllers.GetAllInfluencers())
}
