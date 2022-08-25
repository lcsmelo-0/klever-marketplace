package main

import (
	"klever-marketplace/configs"
	"klever-marketplace/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	router.GET("/influencers", controllers.GetAllInfluencers())

	router.Run()
}
