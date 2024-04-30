package frontend

import (
	"website/src/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)

func Filesystem(c echo.Context) error {
	component := pages.Filesystem()
	return utils.Render(c, 200, component)
}