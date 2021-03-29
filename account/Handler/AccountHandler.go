package Handler

import (
	"github.com/labstack/echo/v4"
	"github.com/superDeano/account/Dao"
	"github.com/superDeano/account/Model"
	"net/http"
	"strconv"
)

func AddAccount(c echo.Context) error {
	user := &Model.Account{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Could not add User\n"+err.Error())
	}
	if Dao.AddUser(user) {
		return c.String(http.StatusOK, "Added user\n")
	} else {
		return c.String(http.StatusBadRequest, "Email or username already in use\n")
	}
}

func AuthenticateUser(c echo.Context) error {
	user := &Model.Account{}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Could not authenticate user\n"+err.Error())
	}
	if Dao.AuthenticateUser(user) {
		return c.JSON(http.StatusOK, true)
	} else {
		return c.JSON(http.StatusBadRequest, false)
	}
}

func DeleteUser(c echo.Context) error {
	switch num := Dao.DeleteUser(c.Param(Dao.AccountUsername)); {
	case num < Dao.Success:
		return c.JSON(http.StatusBadRequest, "Deleted nothing!")
	case num > Dao.Success:
		return c.JSON(http.StatusBadRequest, "Deleted more than one account!")
	default:
		return c.JSON(http.StatusOK, "Deleted!")
	}
}

func GetUser(c echo.Context) error {
	userInfo, err := Dao.GetUserWithUsername(c.Param(Dao.AccountUsername))
	if err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}
	return c.JSON(http.StatusOK, userInfo)

}

func Test(c echo.Context) error {
	if Dao.Db == nil {
		return c.JSON(http.StatusInternalServerError, "The DBPointer is nil")
	}
	if res := Dao.TestDatabaseConnection(); res ==
		Dao.Success {
		return c.JSON(http.StatusOK, "It seems like the app is reachable and database connection is great")
	} else {
		return c.JSON(http.StatusInternalServerError, "It seems like database connection is not too great")
	}

}

func AddShowToWatchlist(c echo.Context) error {
	user := c.Param(Dao.AccountUsername)
	showId, err := strconv.Atoi(c.Param(Dao.ShowId))
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "Error with show id")
	}
	resStatus, errMsg := Dao.AddShowToWatchList(user, showId)
	return c.JSON(resStatus, errMsg)
}

func AddMovieToWatchlist(c echo.Context) error {
	user := c.Param(Dao.AccountUsername)
	movieId, err := strconv.Atoi(c.Param(Dao.MovieId))
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, "Error with movie id")
	}
	resStatus, errMsg := Dao.AddMovieToWatchlist(user, movieId)
	return c.JSON(resStatus, errMsg)

}

func RemoveShowFromWatchlist(c echo.Context) error {
	user := c.Param(Dao.AccountUsername)
	showId, err := strconv.Atoi(c.Param(Dao.ShowId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error with showId!")
	}
	status, res := Dao.RemoveShowFromWatchlist(user, showId)
	return c.JSON(status, res)
}

func RemoveMovieFromWatchlist(c echo.Context) error {
	user := c.Param(Dao.AccountUsername)
	movieId, err := strconv.Atoi(c.Param(Dao.ShowId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Error with movieId!")
	}
	status, res := Dao.RemoveMovieFromWatchlist(user, movieId)
	return c.JSON(status, res)
}

func GetShowWatchlist(c echo.Context) error {
	user := c.Param(Dao.AccountUsername)
	status, res := Dao.GetShowWatchlist(user)
	return c.JSON(status, res)
}

func GetMovieWatchlist(c echo.Context) error {
	user := c.Param(Dao.AccountUsername)
	status, res := Dao.GetMovieWatchlist(user)
	return c.JSON(status, res)
}
