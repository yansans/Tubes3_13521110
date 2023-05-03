package models

type UserMessage struct {
	ChatID       string   `json:"chat_id"`
	ChatName     string   `json:"chat_name"`
	Participants string   `json:"participants"`
	Messages     []string `json:"message"`
}

type UserChat struct {
	ChatID   string `json:"chat_id"`
	ChatName string `json:"chat_name"`
}

type NewUserMessage struct {
	ChatID  string `json:"chat_id"`
	Message string `json:"message"`
}
