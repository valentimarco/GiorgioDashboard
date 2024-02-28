package frontend

import (
	"website/cmd/utils"
	"website/templates"
	"website/templates/components/index"

	"github.com/labstack/echo/v4"
)

var dropdowns = map[string][]string{
	"infos":         {"Dashboard", "Settings"},
	"orchestration": {"Docker", "Kubernetes", "Docker Swarm"},
	"nas":           {"Synology", "QNAP", "FreeNAS"},
	"mailbox":       {"Gmail", "Outlook", "Yahoo"},
}

func IndexPage(c echo.Context) error {
	component := templates.Index()
	return utils.Render(c, 200, component)
}

func Dropdown(c echo.Context) error {
	target := c.Request().Header.Get("hx-target")
	list := dropdowns[target]
	component := index.Dropdown(list)
	return utils.Render(c, 200, component)
}
