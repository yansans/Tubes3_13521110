package routes

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/controllers"
)

func MutexMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	var mutex sync.Mutex

	return func(c echo.Context) error {
		mutex.Lock()
		defer mutex.Unlock()

		err := next(c)
		return err
	}
}

func FrontEndRoute(e *echo.Echo) {
	e.GET("/app", controllers.GetUserChats, MutexMiddleware)
	e.POST("/app", controllers.NewUserChat, MutexMiddleware)
	e.DELETE("/app/:id", controllers.DeleteUserChatID, MutexMiddleware)
	e.DELETE("/app", controllers.DeleteUserChat, MutexMiddleware)
	e.POST("/app/chat", controllers.AddNewUserMessage, MutexMiddleware)
	e.PUT("/app", controllers.RenameUserChat, MutexMiddleware)
}
