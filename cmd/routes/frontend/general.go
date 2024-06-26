package frontend

import (
	"website/cmd/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)

func General(c echo.Context) error {
	component := pages.General()
	return utils.Render(c, 200, component)
}
