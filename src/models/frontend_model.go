package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserMessage struct {
	ChatID       string   `json:"chat_id" validate:"required"`
	ChatName     string   `json:"chat_name" validate:"required"`
	Participants string   `json:"participants" validate:"required"`
	Messages     []string `json:"message" validate:"required"`
}

type NewChat struct {
	ChatID       string             `json:"chat_id"`
	ChatName     string             `json:"chat_name" validate:"required"`
	CreationDate primitive.DateTime `json:"creation_date"`
}

type DeleteChat struct {
	ChatID string `json:"chat_id" validate:"required"`
}

type EditChat struct {
	ChatID   string `json:"chat_id" validate:"required"`
	ChatName string `json:"chat_name" validate:"required"`
}

type NewUserMessage struct {
	ChatID    string `json:"chat_id" validate:"required"`
	Message   string `json:"message" validate:"required"`
	Sender    string `json:"sender"`
	Algorithm string `json:"algorithm"`
}
