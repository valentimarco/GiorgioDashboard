package frontend

import (
	"website/cmd/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	component := pages.Login()
	return utils.Render(c, 200, component)
}
