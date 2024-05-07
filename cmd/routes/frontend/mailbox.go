package frontend

import (
	"website/cmd/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)


func MailBox(c echo.Context) error {
	component := pages.MailBox()
	return utils.Render(c, 200, component)
}