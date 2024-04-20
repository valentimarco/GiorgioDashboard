package main

import (
	"website/cmd/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))
	e.Static("/css", "css")
	e.Static("/public", "public")
	e.Static("/static", "static")
	routes.IndexRoutes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
