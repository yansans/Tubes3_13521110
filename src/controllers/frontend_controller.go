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
	var chat models.UserChat
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}
	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}
	chat.ChatID = primitive.NewObjectID().Hex()

	uri := "http://localhost:6969/chat"

	chatBody, err := json.Marshal(&chat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "marshal", Data: &echo.Map{"data": err.Error()}})
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, uri, bytes.NewBuffer(chatBody))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "new", Data: &echo.Map{"data": err.Error()}})
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "do", Data: &echo.Map{"data": err.Error()}})
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {

		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "read", Data: &echo.Map{"data": err.Error()}})
	}

	var respUser responses.UserResponse
	if err := json.Unmarshal(body, &respUser); err != nil {

		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "unmarshal", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": respUser.Data}})
}

func DeleteUserChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var chat models.UserChat
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}
	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	uri := "http://localhost:6969/chat/" + chat.ChatID
	println(uri)
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "request", Data: &echo.Map{"data": err.Error()}})
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "request", Data: &echo.Map{"data": err.Error()}})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "status", Data: &echo.Map{"data": resp.StatusCode}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: nil})
}
