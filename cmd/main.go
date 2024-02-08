package main

import (
	"website/templates"
	"website/templates/components"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	count := Count{Count: 0}

	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
		return Render(c, 200, component)
	})

	e.POST("/count", func(c echo.Context) error {
		count.Count++
		component := components.Count(count.Count)
		return Render(c, 200, component)
	})

	e.Static("/css", "css")
	e.Static("/static", "static")
	e.Static("/fonts", "fonts")
	e.Logger.Fatal(e.Start(":3000"))
}
