package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/configs"
	"github.com/yansans/Tubes3_13521110/src/models"
	"github.com/yansans/Tubes3_13521110/src/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var chatLogColection *mongo.Collection = configs.GetCollection(configs.DB, "chatlogs")

func CreateChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var chat models.ChatLog
	defer cancel()

	//validate the request body
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	if chat.Messages == nil {
		chat.Messages = make([]models.Message, 0)
	}

	newChat := models.ChatLog{
		ID:           primitive.NewObjectID(),
		ChatName:     chat.ChatName,
		Participants: "user",
		Messages:     chat.Messages,
		TotalMessage: len(chat.Messages),
		CreationDate: primitive.NewDateTimeFromTime(time.Now()),
	}

	result, err := chatLogColection.InsertOne(ctx, newChat)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func GetAllChats(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var chats []models.ChatLog
	defer cancel()

	opts := options.Find().SetSort(bson.M{"creation_date": 1})
	cursor, err := chatLogColection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err = cursor.All(ctx, &chats); err != nil {
		return e.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if chats == nil {
		chats = make([]models.ChatLog, 0)
	}

	return e.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": chats}})
}

func GetChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var chat models.ChatLog
	defer cancel()

	chatIDHex := c.Param("id")
	chatID, err := primitive.ObjectIDFromHex(chatIDHex)
	if err != nil {
		return err
	}

	if err := chatLogColection.FindOne(ctx, bson.M{"_id": chatID}).Decode(&chat); err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": chat}})
}

func SendMessage(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var chat models.Message
	defer cancel()

	chatIDHex := c.Param("id")
	chatID, err := primitive.ObjectIDFromHex(chatIDHex)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "convert", Data: &echo.Map{"data": err.Error()}})
	}

	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "validate", Data: &echo.Map{"data": validationErr.Error()}})
	}

	chat.ID = primitive.NewObjectID()
	chat.Timestamp = primitive.NewDateTimeFromTime(time.Now())

	result, err := chatLogColection.UpdateOne(ctx, bson.M{"_id": chatID}, bson.M{"$push": bson.M{"messages": chat}, "$inc": bson.M{"total_message": 1}})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "update", Data: &echo.Map{"data": err.Error()}})
	}

	if result.ModifiedCount == 0 {
		return c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "not found", Data: &echo.Map{"data": "chat not found"}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}

func DeleteChat(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatIDHex := e.Param("id")
	chatID, err := primitive.ObjectIDFromHex(chatIDHex)

	if err != nil {
		return e.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "convert", Data: &echo.Map{"data": err.Error()}})
	}

	result, err := chatLogColection.DeleteOne(ctx, bson.M{"_id": chatID})

	if err != nil {
		return e.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount == 0 {
		return e.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "chat not found"}})
	}

	return e.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}

func RenameChat(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chatIDHex := c.Param("id")
	chatID, err := primitive.ObjectIDFromHex(chatIDHex)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "convert", Data: &echo.Map{"data": err.Error()}})
	}

	var chat models.ChatLog
	if err := c.Bind(&chat); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "bind", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&chat); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "validate", Data: &echo.Map{"data": validationErr.Error()}})
	}

	filter := bson.M{"_id": chatID}
	update := bson.M{"$set": bson.M{"chat_name": chat.ChatName}}

	result, err := chatLogColection.UpdateOne(ctx, filter, update)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "update", Data: &echo.Map{"data": err.Error()}})
	}

	if result.ModifiedCount == 0 {
		return c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "not found", Data: &echo.Map{"data": "chat not found"}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}
