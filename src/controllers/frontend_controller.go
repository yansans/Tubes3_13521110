package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/models"
	"github.com/yansans/Tubes3_13521110/src/responses"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserChats(c echo.Context) error {
	resp, err := http.Get("http://localhost:6969/chat")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "get", Data: &echo.Map{"data": err.Error()}})
	}

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "status", Data: &echo.Map{"data": resp.StatusCode}})
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "io", Data: &echo.Map{"data": err.Error()}})
	}

	var respUser responses.UserResponse
	if err := json.Unmarshal(body, &respUser); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "unmarshal", Data: &echo.Map{"data": err.Error()}})
	}

	result := respUser.Data
	chatlog := (*result)["data"]

	chatlogs := make([]models.UserMessage, 0)

	for _, chat := range chatlog.([]interface{}) {
		chatMap := chat.(map[string]interface{})
		messages := make([]string, 0)

		for _, message := range chatMap["messages"].([]interface{}) {
			messageMap := message.(map[string]interface{})
			messageString := messageMap["message"].(string)
			messages = append(messages, messageString)
			// id := messageMap["id"].(string)
			// timestampStr := messageMap["timestamp"].(string)
			// timestamp, err := primitive.ParseDateTime(timestampStr)
			// ID:      objectID,
			// Sender:  messageMap["sender"].(string),

		}
		ID := chatMap["id"].(string)
		chatName := chatMap["chat_name"].(string)
		participants := chatMap["participants"].(string)

		newUserMessage := models.UserMessage{
			ChatID:       ID,
			ChatName:     chatName,
			Participants: participants,
			Messages:     messages,
			// Participants: chatMap["participants"].(string),
			// TotalMessage: int(chatMap["total_message"].(float64)),
			// CreationDate: chatMap["creation_date"].(string),
		}

		chatlogs = append(chatlogs, newUserMessage)
	}

	return c.JSON(http.StatusOK, responses.ChatFrontendResponse{Status: http.StatusOK, Message: "success", Data: chatlogs})
}

func NewUserChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var chat models.NewChat
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	chatID := primitive.NewObjectID()
	chat.ChatID = chatID.Hex()
	chat.CreationDate = primitive.NewDateTimeFromTime(time.Now())

	chatData, err := json.Marshal(chat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "marshal", Data: &echo.Map{"data": err.Error()}})
	}

	url := "http://localhost:6969/chat"

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(chatData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "new request", Data: &echo.Map{"data": err.Error()}})
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "client", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "read all", Data: &echo.Map{"data": err.Error()}})
	}

	if resp.StatusCode != http.StatusCreated {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "status", Data: &echo.Map{"data": string(body)}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: nil})
}

func DeleteUserChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	var chat models.DeleteChat
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}
	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	url := "http://localhost:6969/chat/" + chat.ChatID
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "new request", Data: &echo.Map{"data": err.Error()}})
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "client", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "read all", Data: &echo.Map{"data": err.Error()}})
	}

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "status", Data: &echo.Map{"data": string(body)}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: nil})
}

func RenameUserChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var chat models.EditChat

	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}
	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	url := "http://localhost:6969/chat/" + chat.ChatID
	chatData, err := json.Marshal(chat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "marshal", Data: &echo.Map{"data": err.Error()}})
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, bytes.NewBuffer(chatData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "new request", Data: &echo.Map{"data": err.Error()}})
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "client", Data: &echo.Map{"data": err.Error()}})
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "read all", Data: &echo.Map{"data": err.Error()}})
	}

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "status", Data: &echo.Map{"data": string(body)}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: nil})
}

func AddNewUserMessage(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var message models.NewUserMessage
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})

	}
	if validationErr := validate.Struct(&message); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	message.Sender = "user"
	url := "http://localhost:6969/chat/" + message.ChatID
	messageData, err := json.Marshal(message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "marshal", Data: &echo.Map{"data": err.Error()}})
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(messageData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "new request", Data: &echo.Map{"data": err.Error()}})
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "client", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "read all", Data: &echo.Map{"data": err.Error()}})
	}

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "status", Data: &echo.Map{"data": string(body)}})
	}

	// fungsi respon disini
	responsesString := message.Message
	println(responsesString)

	var botMessage models.NewUserMessage
	botMessage.Sender = "bot"
	botMessage.ChatID = message.ChatID
	botMessage.Message = "responsesString"
	botMessageData, err := json.Marshal(botMessage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "bot marshal", Data: &echo.Map{"data": err.Error()}})
	}

	req, err = http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(botMessageData))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "bot new request", Data: &echo.Map{"data": err.Error()}})
	}

	req.Header.Set("Content-Type", "application/json")
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "bot client", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "bot read all", Data: &echo.Map{"data": err.Error()}})
	}

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "bot status", Data: &echo.Map{"data": string(body)}})
	}

	return c.JSON(http.StatusOK, responses.ChatFrontendResponse{Status: http.StatusOK, Message: "success", Data: responsesString})
}
