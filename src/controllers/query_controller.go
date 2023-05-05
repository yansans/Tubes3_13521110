package controllers

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/configs"
	"github.com/yansans/Tubes3_13521110/src/models"
	"github.com/yansans/Tubes3_13521110/src/responses"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var queryCollection *mongo.Collection = configs.GetCollection(configs.DB, "queries")

func CreateQuery(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var query models.Query
	defer cancel()

	//validate the request body
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&query); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newQuery := models.Query{
		ID:       primitive.NewObjectID(),
		Question: query.Question,
		Answer:   query.Answer,
	}

	result, err := queryCollection.InsertOne(ctx, newQuery)

	if err != nil {
		if mongo.IsDuplicateKeyError(err) {

			update := bson.M{"$set": bson.M{"answer": newQuery.Answer}}

			result, err := queryCollection.UpdateOne(ctx, bson.M{"question": newQuery.Question}, update)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
			}

			return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success replaced", Data: &echo.Map{"data": result}})

		} else {

			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}
	}
	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func DeleteQuery(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	id := c.Param("queryId")
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(id)

	result, err := queryCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "User with specified _id not found!"}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}

func DeleteQueryQuestion(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var query models.Query
	defer cancel()

	//validate the request body
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&query); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	question := query.Question

	result, err := queryCollection.DeleteOne(ctx, bson.M{"question": question})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "error", Data: &echo.Map{"data": "User with specified _id not found!"}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})
}

func GetAllQueries(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var queries []models.Query
	defer cancel()

	results, err := queryCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuery models.Query
		if err = results.Decode(&singleQuery); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		queries = append(queries, singleQuery)
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": queries}})
}

func GetAnswer(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var query models.Query
	defer cancel()

	//validate the request body
	if err := c.Bind(&query); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&query); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	question := query.Question

	err := queryCollection.FindOne(ctx, bson.M{"question": question}).Decode(&query)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": query.Answer}})
}

func GetQuestionList() []string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var questions []string

	results, err := queryCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuery models.Query
		if err = results.Decode(&singleQuery); err != nil {
			return nil
		}

		questions = append(questions, strings.ToLower(singleQuery.Question))
	}

	return questions
}

func GetQuestionMap() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	questions := make(map[string]string)

	results, err := queryCollection.Find(ctx, bson.M{})

	if err != nil {
		return nil
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleQuery models.Query
		if err = results.Decode(&singleQuery); err != nil {
			return nil
		}

		questions[strings.ToLower(singleQuery.Question)] = singleQuery.Answer
	}

	return questions
}
