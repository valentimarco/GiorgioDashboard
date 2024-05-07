package frontend


import (
	"website/cmd/utils"
	"website/templates/pages"

	"github.com/labstack/echo/v4"
)

// Docker renders the Docker page
func Docker(c echo.Context) error {
	component := pages.Docker()
	return utils.Render(c, 200, component)
}