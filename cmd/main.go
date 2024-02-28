package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"website/cmd/routes"
	
)

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Static("/fonts", "fonts")
	routes.IndexRoutes(e)


	


	// count := Count{Count: 0}

	// e.GET("/", func(c echo.Context) error {
	// 	component := templates.Index()
	// 	return utils.Render(c, 200, component)
	// })

	// e.POST("/count", func(c echo.Context) error {
	// 	count.Count++
	// 	component := components.Count(count.Count)
	// 	return utils.Render(c, 200, component)
	// })

	// e.POST("/dropdown", func(c echo.Context) error {
	// 	list := []string{"Home", "About", "Contact"}
	// 	component := index.Dropdown(list)
	// 	return utils.Render(c, 200, component)
	// })

	e.Logger.Fatal(e.Start(":3000"))
}


