package Handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddAccount(c echo.Context) error {
	return c.String(http.StatusOK, "Trying to add user, well did not fucking work")
}

func AuthenticateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, false)
}

func Test(c echo.Context) error {
	return c.JSON(http.StatusOK, "It seems like the app is reachable")
}