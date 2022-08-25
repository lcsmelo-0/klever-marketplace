package controllers

import (
	"context"
	"klever-marketplace/configs"
	"klever-marketplace/models"
	"klever-marketplace/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "influencers")
var validate = validator.New()

func CreateInfluencer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var influencer models.Influencer
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&influencer); err != nil {
			c.JSON(http.StatusBadRequest, responses.InfluencerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&influencer); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.InfluencerResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newInfluencer := models.Influencer{
			Id:                 primitive.NewObjectID(),
			Name:               influencer.Name,
			InstagramProfile:   influencer.InstagramProfile,
			InstagramFollowers: influencer.InstagramFollowers,
			ProfileDescription: influencer.ProfileDescription,
		}

		result, err := userCollection.InsertOne(ctx, newInfluencer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InfluencerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.InfluencerResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}
