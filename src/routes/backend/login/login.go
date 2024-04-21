package login

import (
	"github.com/labstack/echo/v4"
)


func Login(c echo.Context) error {
	println(c.Request().FormValue("username"))
	return c.String(200, "Login")
}
