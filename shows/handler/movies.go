package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"shows/model"
	"shows/variable"
	"log"
)


func SearchMovie(c echo.Context) error {
	keywords := c.Param(variable.Keywords)
	options, err := getSearchOptions(c)
	res, err := tmdbApi.SearchMovie(keywords, options)
	if err != nil {
		log.Println("Error when trying to search for movies")
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, res)
}


func GetTrendingMovies(c echo.Context) error {
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	res, err := GetTrending(string(model.Movie), pageNumber)
	if err != nil {
		log.Println("Error when getting Trending Movies")
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, res)
}

func GetMovieWatchProvider(c echo.Context) error {
	movieId := c.Param(variable.Id)
	url := fmt.Sprintf("%smovie/%v/watch/providers%s", variable.BaseUrl,movieId, getApiAuth())
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error when searching for Show Providers")
	}
	var providers = &model.WatchProvider{}
	if err := json.NewDecoder(res.Body).Decode(&providers); err != nil {
		log.Println("Error when decoding Show Providers")
		return c.JSON(http.StatusExpectationFailed, providers)
	}
	return c.JSON(http.StatusOK, providers)
}