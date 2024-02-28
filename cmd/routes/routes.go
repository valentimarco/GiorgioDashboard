package routes

import (
	"website/cmd/routes/frontend"

	"github.com/labstack/echo/v4"
)

func IndexRoutes(e *echo.Echo) {
	e.GET("/", frontend.IndexPage)
	e.POST("/dropdown", frontend.Dropdown)
}
