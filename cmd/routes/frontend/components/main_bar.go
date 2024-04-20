package components

import ( 
	"website/cmd/utils"
	"website/templates/components/index"

	"github.com/labstack/echo/v4"
)

func MainBar(c echo.Context) error {
	title := c.QueryParam("page")
	component := index.MainBar(title)
	return utils.Render(c, 200, component)
}

func Avatar(c echo.Context) error {
	component := index.Avatar()
	return utils.Render(c, 200, component)
}
