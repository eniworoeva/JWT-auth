package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID					primitive.ObjectID	`bson:"_id"`
	FirstName			*string				`json:"firstname" validate:"required, min=3, max=20"`
	LastName			*string				`json:"lastname" validate:"required, min=3, max=20"`
	Password			*string				`json:"password" validate:"required, min=5"`
	Email				*string				`json:"email" validate:"required, email"`
	Phone				*string				`json:"phone" validate:"required"`
	Token				*string				`json:"token"`
	User_type			*string				`json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
	Refresh_token		*string				`json:"refresh_token"`
	Created_at			time.Time			`json:"created_at"`
	Updated_at			time.Time			`json:"updated_at"`
	User_id				string				`json:"user_id"`
}