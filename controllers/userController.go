package controllers

import (
	"github.com/eniworoeva/JWT-auth/database"
	"github.com/go-playground/validator/v10"
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

func GetUsers()

func GetUser()