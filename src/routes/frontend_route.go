package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/controllers"
)

func FrontEndRoute(e *echo.Echo) {
	e.GET("/app", controllers.GetUserChats)
	e.POST("/app", controllers.NewUserChat)
	// e.DELETE("/app/:id", controllers.DeleteUserChatID)
	e.DELETE("/app", controllers.DeleteUserChat)
	e.POST("/app/chat", controllers.AddNewUserMessage)
	e.PUT("/app/chat", controllers.RenameUserChat)
}
