package Handler

import (
	"github.com/labstack/echo/v4"
	"github.com/superDeano/account/Dao"
	"github.com/superDeano/account/Model"
	"net/http"
)

func AddAccount(c echo.Context) error {
	user := &Model.Account{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Could not add User\n"+err.Error())
	}
	Dao.AddUser(user)
	return c.String(http.StatusOK, "Added user\n")
}

func AuthenticateUser(c echo.Context) error {
	return c.JSON(http.StatusOK, true)
}

func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, false)
}

func Test(c echo.Context) error {
	if res := Dao.TestDatabaseConnection(); res ==
		Dao.Success {
		return c.JSON(http.StatusOK, "It seems like the app is reachable and database connection is great")
	} else {
		return c.JSON(http.StatusInternalServerError, "It seems like database connection is not too great")
	}

}
