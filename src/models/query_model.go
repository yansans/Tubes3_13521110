package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Query struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Question string             `bson:"question" json:"question" validate:"required"`
	Answer   string             `bson:"answer" json:"answer"`
}
