package routes

import (
	"website/cmd/models"
	"website/cmd/routes/frontend"
	"website/cmd/routes/frontend/components"
	"website/cmd/utils"
	"website/templates/components/charts"

	"github.com/labstack/echo/v4"
)

func IndexRoutes(e *echo.Echo) {
	e.GET("/", frontend.IndexPage)
	e.GET("/mainbar", components.MainBar)
	e.POST("/avatar", components.Avatar)
	e.GET("/setting", frontend.Setting)
	e.GET("/general", frontend.General)
	e.GET("/chart", func(c echo.Context) error {
		// Generate the data for the chart
		data := charts.Chart(models.Votes{Red: 1, Blue: 2, Yellow: 3, Green: 4, Purple: 5, Orange: 6})
		// Render the chart
		return utils.Render(c, 200, data)
	})
}
