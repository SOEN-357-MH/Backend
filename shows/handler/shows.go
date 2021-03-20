package handler

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shows/model"
	"shows/variable"
	"strconv"
)

func GetHealth(c echo.Context) error {
	configUrl := variable.Base_url + "configuration" + getApiAuth()
	resp, err := http.Get(configUrl)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Cannot reach External API Source")
	}
	var config model.Configuration
	if err = json.NewDecoder(resp.Body).Decode(&config); err != nil {
		log.Fatal("Error when decoding information about the Configurations")
	}

	defer resp.Body.Close()

	return c.String(http.StatusOK, "App seems reachable!")
}

func ConfigureSomeKeyVariables() {
	configUrl := variable.Base_url + "configuration" + getApiAuth()
	resp, err := http.Get(configUrl)
	if err != nil {
		log.Fatal("Cannot reach External API Source")
	}
	var config model.Configuration
	if err = json.NewDecoder(resp.Body).Decode(&config); err != nil {
		log.Fatal("Error when decoding information about the Configurations")
	}
	go AssignVariables(&config)
	defer resp.Body.Close()
}

func AssignVariables(config *model.Configuration) {
	variable.Image_url = *config.Images.BaseURL
	variable.Image_size = config.Images.ProfileSizes[len(config.Images.ProfileSizes)-1]
	variable.ChangeKeys = make(map[string]string)
	for _, value := range config.ChangeKeys {
		variable.ChangeKeys[value] = value
	}
}

func GetTrendingMovies(c echo.Context) error {
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	if err != nil || pageNumber < 1 || pageNumber > 1000 {
		pageNumber = 1
	}
	res, err := GetTrending(string(model.Movie), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func GetTrendingShows(c echo.Context) error {
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	if err != nil || pageNumber < 1 || pageNumber > 1000 {
		pageNumber = 1
	}
	res, err := GetTrending(string(model.Tv), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func GetTrendingPersons(c echo.Context) error {
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	if err != nil || pageNumber < 1 || pageNumber > 1000 {
		pageNumber = 1
	}
	res, err := GetTrending(string(model.Person), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func GetTrendingAll(c echo.Context) error {
	pageNumber, err := strconv.Atoi(c.Param("pageNumber"))
	if err != nil || pageNumber < 1 || pageNumber > 1000 {
		pageNumber = 1
	}
	res, err := GetTrending(string(model.All), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func GetTrending(mediaType string, pageNumber int) (*model.Result, error) {
	url := variable.Base_url + "trending/" + mediaType + "/" + model.Day + getApiAuth() + "&pageNumber=" + string(rune(pageNumber))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	results := &model.Result{}
	if err = json.NewDecoder(res.Body).Decode(results); err != nil {
		return nil, err
	}
	return results, nil
}

func GetImageSize(c echo.Context) error {
	return c.String(http.StatusOK, variable.Image_size)
}

func GetImageBaseUrl(c echo.Context) error {
	return c.String(http.StatusOK, variable.Image_url)
}

func GetChangeKeys(c echo.Context) error {
	return c.JSON(http.StatusOK, variable.ChangeKeys)
}

func getApiAuth() string {
	return "?api_key=" + variable.Api_key
}

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}
