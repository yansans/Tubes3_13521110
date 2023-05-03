package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ChatLog struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChatName     string             `bson:"chat_name" json:"chat_name"`
	Participants string             `bson:"participants" json:"participants"`
	Messages     []Message          `bson:"messages" json:"messages"`
	TotalMessage int                `bson:"total_message" json:"total_message"`
	CreationDate primitive.DateTime `bson:"creation_date" json:"creation_date"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Sender    string             `bson:"sender" json:"sender"`
	Message   string             `bson:"message" json:"message"`
	Timestamp primitive.DateTime `bson:"timestamp" json:"timestamp"`
}
