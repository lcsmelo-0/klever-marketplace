package main

import (
	"klever-marketplace/configs"
	"klever-marketplace/routes"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     cors.DefaultConfig().AllowMethods,
		AllowHeaders:     cors.DefaultConfig().AllowHeaders,
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	configs.ConnectDB()

	routes.UserRoute(router)
	routes.OrderRoute(router)

	httpPort := ":8080"
	if os.Getenv("PORT") != "" {
		httpPort = ":" + os.Getenv("PORT")
	}

	router.Run(httpPort)
}
