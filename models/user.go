package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id            bson.ObjectId `json:"id" bson:"_id"`
    Name    	  string            `json:"name" bson:"name"`
	Email         string            `json:"email" bson:"email"`
    Password      string            `json:"password" bson:"password"`
	Title         string            `json:"title" bson:"title"`
	Birthdate	  string		    `json:"birthdate" bson:"birthdate"`
    CreatedAt     time.Time         `json:"created_at" bson:"created_at"`
    UpdatedAt     time.Time         `json:"updated_at" bson:"updated_at"`
	User_type     string            `json:"user_type" bson:"user_type" validate:"required,eq=ADMIN|eq=USER""`
}
