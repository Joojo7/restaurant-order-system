package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Food is the model that governs all notes objects retrived or inserted into the DB
type Food struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `json:"name"`
	Price     float32            `json:"price"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	FoodID    string             `json:"food_id"`
	MenuID    string             `json:"menu_id"`
}
