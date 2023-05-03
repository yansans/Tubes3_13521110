package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/controllers"
)

func FrontEndRoute(e *echo.Echo) {
	e.GET("/app", controllers.GetUserChats)
	// e.POST("/app", controllers.CreateChat)
	// e.DELETE("/app", controllers.DeleteChat)
}
