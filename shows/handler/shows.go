package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ryanbradynd05/go-tmdb"
	"log"
	"net/http"
	"shows/model"
	"shows/variable"
	"strconv"
)

var tmdbApi *tmdb.TMDb

func GetHealth(c echo.Context) error {
	configUrl := fmt.Sprintf("%sconfiguration/%s", variable.BaseUrl, getApiAuth())
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

func setUpTmdbLib() {
	config := tmdb.Config{
		APIKey:   variable.ApiKey,
		Proxies:  nil,
		UseProxy: false,
	}
	tmdbApi = tmdb.Init(config)
}

func ConfigureSomeKeyVariables() {
	configUrl := variable.BaseUrl + "configuration" + getApiAuth()
	resp, err := http.Get(configUrl)
	if err != nil {
		log.Fatal("Cannot reach External API Source")
	}
	var config model.Configuration
	if err = json.NewDecoder(resp.Body).Decode(&config); err != nil {
		log.Fatal("Error when decoding information about the Configurations")
	}
	go AssignVariables(&config)
	go setUpTmdbLib()
	defer resp.Body.Close()
}

func AssignVariables(config *model.Configuration) {
	variable.ImageUrl = *config.Images.BaseURL
	variable.ImageSize = config.Images.ProfileSizes[len(config.Images.ProfileSizes)-1]
	variable.ChangeKeys = make(map[string]string)
	for _, value := range config.ChangeKeys {
		variable.ChangeKeys[value] = value
	}
}

func GetTrendingShows(c echo.Context) error {
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	res, err := GetTrending(string(model.Tv), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func getPageNumber(page string) (string, error) {
	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 || pageNumber > 1000 {
		pageNumber = 1
	}
	return strconv.Itoa(pageNumber), err
}

func GetTrendingAll(c echo.Context) error {
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	res, err := GetTrending(string(model.All), pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func GetTrending(mediaType string, pageNumber string) (*model.Result, error) {
	url := variable.BaseUrl + "trending/" + mediaType + "/" + model.Day + getApiAuth() + "&pageNumber=" + pageNumber
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
	return c.String(http.StatusOK, variable.ImageSize)
}

func GetImageBaseUrl(c echo.Context) error {
	return c.String(http.StatusOK, variable.ImageUrl)
}

func GetChangeKeys(c echo.Context) error {
	return c.JSON(http.StatusOK, variable.ChangeKeys)
}

func getApiAuth() string {
	return "?api_key=" + variable.ApiKey
}

func getSearchOptions(c echo.Context) (map[string]string, error) {
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	if err != nil {
		log.Println("Error when parsing page number, page number set to default")
	}
	options := make(map[string]string)
	options["language"] = string(model.En)
	options["page"] = pageNumber
	return options, err
}

func SearchShows(c echo.Context) error {
	keywords := c.Param(variable.Keywords)
	options, err := getSearchOptions(c)
	if err != nil {
		log.Fatal("Error during parsing of options, setting variables to default")
	}
	res, err := tmdbApi.SearchTv(keywords, options)
	if err != nil {
		log.Println("Error when searching for Shows")
		return c.JSON(http.StatusInternalServerError, res)
	}
	return c.JSON(http.StatusOK, res)

}

func GetShowWatchProvider(c echo.Context) error {
	showId := c.Param(variable.Id)
	url := fmt.Sprintf("%vtv/%v/watch/providers%s", variable.BaseUrl,showId, getApiAuth())
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

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}
