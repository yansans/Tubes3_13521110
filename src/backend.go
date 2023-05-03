package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yansans/Tubes3_13521110/configs"
	"github.com/yansans/Tubes3_13521110/src/routes"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	configs.ConnectDB()

	routes.UserRoute(e)
	routes.QueryRoute(e)
	routes.ChatLogRoute(e)

	routes.FrontEndRoute(e)

	e.Logger.Fatal(e.Start("localhost:6969"))
}
