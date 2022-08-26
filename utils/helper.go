package utils

import (
	"os"

	"github.com/gin-gonic/gin"
)

// SendJSONResponse formats response to JSON
func SendJSONResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// SendJSONError formats error response to JSON
func SendJSONError(c *gin.Context, statusCode int, err error) {
	SendJSONResponse(c, statusCode, gin.H{
		"error": err.Error(),
	})
}

func GetMongoURI() string {
	return os.Getenv("MONGOURI")
}
