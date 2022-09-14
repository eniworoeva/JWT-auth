package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/genproto/googleapis/type/phone_number"
)

type User struct {
	ID					primitive.ObjectID	 `bson:"_id"`
	FirstName			*string				`json:"firstname" validate:"required, min=3, max=20"`
	LastName			*string				`json:"lastname" validate:"required, min=3, max=20"`
	Password			*string
	Email				*string
	Phone				*string
	Token				*string
	User_type			*string
	Refresh_token		*string
	Created_at			time.Time
	Updated_at			time.Time
	User_id				string
}