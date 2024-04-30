package frontend
import (
	"website/src/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)


func Test(c echo.Context) error {
	component := pages.Test()
	return utils.Render(c, 200, component)
}