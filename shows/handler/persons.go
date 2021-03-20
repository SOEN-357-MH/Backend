package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"shows/model"
	"shows/variable"
	"log"
)

func GetTrendingPersons(c echo.Context) error {
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	res, err := GetTrending(string(model.Person), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}
