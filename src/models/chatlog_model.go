package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatLog struct {
	ID           primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Participants []primitive.ObjectID `bson:"participants" json:"participants" validate:"required"`
	Messages     []Message            `bson:"messages" json:"messages"`
	TotalMessage int                  `bson:"total_message" json:"total_message"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender    primitive.ObjectID `bson:"sender" json:"sender" validate:"required"`
	Message   string             `bson:"message" json:"message" validate:"required"`
	Timestamp primitive.DateTime `bson:"timestamp" json:"timestamp"`
}
