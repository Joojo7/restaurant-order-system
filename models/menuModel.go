package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Menu is the model that governs all notes objects retrived or inserted into the DB
type Menu struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"`
	Category   string             `json:"category"`
	Start_Date time.Time          `json:"start_date"`
	End_Date   time.Time          `json:"end_date"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	MenuID     string             `json:"food_id"`
}
