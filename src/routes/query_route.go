package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/src/controllers"
)

func QueryRoute(e *echo.Echo) {
	e.GET("/query", controllers.GetAllQueries)
	e.POST("/query", controllers.CreateQuery)
	e.DELETE("/query", controllers.DeleteQueryQuestion)
	e.DELETE("/query/:queryId", controllers.DeleteQuery)
	e.GET("/query/ask", controllers.GetAnswer)
}
