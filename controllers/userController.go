package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/eniworoeva/JWT-auth/database"
	"github.com/eniworoeva/JWT-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "User")

var validate = validator.New()

func HashPassword()  {
	
}

func VerifyPassword()  {
	
}

func SignUp()  {
	
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Params("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error( )})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var user models.User
		userCollection.FindOne(ctx, bson.M{"user_id":userId})
	}
}

func GetUser()