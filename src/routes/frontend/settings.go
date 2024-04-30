package frontend

import (
	"website/src/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)

func Settings(c echo.Context) error {
	component := pages.Settings()
	return utils.Render(c, 200, component)
}
