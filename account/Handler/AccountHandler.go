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
	if Dao.AddUser(user) {
		return c.String(http.StatusOK, "Added user\n")
	} else {
		return c.String(http.StatusBadRequest, "Did not add user\n")
	}
}

func AuthenticateUser(c echo.Context) error {
	user := &Model.Account{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Could not authenticate user\n"+err.Error())
	}
	return c.JSON(http.StatusOK, Dao.AuthenticateUser(user))
}

func DeleteUser(c echo.Context) error {
	return c.JSON(http.StatusOK, Dao.DeleteUser(c.Param(Dao.AccountUsername)))
}

func GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, Dao.GetUserWithUsername(c.Param(Dao.AccountUsername)))

}

func Test(c echo.Context) error {
	if res := Dao.TestDatabaseConnection(); res ==
		Dao.Success {
		return c.JSON(http.StatusOK, "It seems like the app is reachable and database connection is great")
	} else {
		return c.JSON(http.StatusInternalServerError, "It seems like database connection is not too great")
	}

}
