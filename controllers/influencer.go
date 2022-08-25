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
	"gopkg.in/mgo.v2/bson"
)

var influencerCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateInfluencer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var influencer models.Influencer
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&influencer); err != nil {
			c.JSON(http.StatusBadRequest, responses.InfluencerResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&influencer); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.InfluencerResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
			return
		}

		newInfluencer := models.Influencer{
			Id:                 primitive.NewObjectID(),
			Name:               influencer.Name,
			InstagramProfile:   influencer.InstagramProfile,
			InstagramFollowers: influencer.InstagramFollowers,
			ProfileDescription: influencer.ProfileDescription,
		}

		result, err := influencerCollection.InsertOne(ctx, newInfluencer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InfluencerResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusCreated, responses.InfluencerResponse{Status: http.StatusCreated, Message: "success", Data: result})
	}
}

func GetInfluencerByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		id := c.Param("influencerId")
		var influencer models.Influencer
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(id)

		err := influencerCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&influencer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InfluencerResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
			return
		}

		c.JSON(http.StatusOK, responses.InfluencerResponse{Status: http.StatusOK, Message: "success", Data: influencer})
	}
}

func GetAllInfluencers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var influencers []models.Influencer
		defer cancel()

		results, err := influencerCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.InfluencerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleInfluencer models.Influencer
			if err = results.Decode(&singleInfluencer); err != nil {
				c.JSON(http.StatusInternalServerError, responses.InfluencerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			influencers = append(influencers, singleInfluencer)
		}

		c.JSON(http.StatusOK,
			responses.InfluencerResponse{Status: http.StatusOK, Message: "success", Data: influencers},
		)
	}
}
