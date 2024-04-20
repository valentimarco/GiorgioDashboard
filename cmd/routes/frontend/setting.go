package frontend

import (
	"website/cmd/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)


func Setting(c echo.Context) error {
	component := pages.Setting()
	return utils.Render(c, 200, component)
}