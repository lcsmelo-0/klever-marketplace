package main

import (
	"klever-marketplace/configs"
	"klever-marketplace/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("localhost:6000")
}
