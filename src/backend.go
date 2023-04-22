package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yansans/Tubes3_13521110/configs"
	"github.com/yansans/Tubes3_13521110/src/routes"
)

func main() {
	e := echo.New()

	configs.ConnectDB()

	routes.UserRoute(e)
	routes.QueryRoute(e)
	routes.ChatLogRoute(e)

	e.Logger.Fatal(e.Start("localhost:6969"))
}