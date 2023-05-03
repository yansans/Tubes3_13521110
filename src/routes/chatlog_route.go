package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/controllers"
)

func ChatLogRoute(e *echo.Echo) {
	e.GET("/chat", controllers.GetAllChats)
	e.POST("/chat", controllers.CreateChat)
	e.GET("/chat/:id", controllers.GetChat)
	e.POST("/chat/:id", controllers.SendMessage)
	e.DELETE("/chat/:id", controllers.DeleteChat)
	e.PUT("/chat/:id", controllers.RenameChat)
}
