package login

import (
	"website/cmd/models"
	"website/cmd/services"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	user := new(models.UserDTO)
	if err := c.Bind(user); err != nil {
		return c.String(400, "Bad request...")
	}
	appstate := c.(*models.AppContext)
	statusCode := services.InsertUser(user, appstate.DB)
	return c.String(statusCode, "Login done")
}

func Login(c echo.Context) error {
	user := new(models.UserDTO)
	if err := c.Bind(user); err != nil {
		return c.String(400, "Bad request...")
	}
	appstate := c.(*models.AppContext)
	statusCode := services.FindUser(user, appstate.DB)
	return c.String(statusCode, "Login done")
}