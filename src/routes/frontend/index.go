package frontend

import (
	"website/src/utils"
	"website/templates"

	"github.com/labstack/echo/v4"
)

func IndexPage(c echo.Context) error {
	component := templates.Index()
	return utils.Render(c, 200, component)
}
