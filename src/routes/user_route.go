package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/controllers"
)

func UserRoute(e *echo.Echo) {
	e.GET("/user", controllers.GetAllUsers)
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetAUser)
	e.PUT("/user/:userId", controllers.EditAUser)
	e.DELETE("/user/:userId", controllers.DeleteAUser)
}
