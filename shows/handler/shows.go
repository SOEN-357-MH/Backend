package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ryanbradynd05/go-tmdb"
	"log"
	"net/http"
	"net/url"
	"shows/model"
	"shows/variable"
	"strconv"
)

var tmdbApi *tmdb.TMDb
var movieGenres *tmdb.Genre
var showsGenres *tmdb.Genre
var config model.Configuration
var Movie = 0
var Show = 1

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
	if variable.ApiKey == "" {
		return c.String(http.StatusFound, "API Key is empty")
	}

	return c.String(http.StatusOK, "App seems reachable!")
}

func setUpTmdbLib() {
	for variable.ApiKey == "" {
		log.Println("Waiting as API Key is empty")
	}
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

	if err = json.NewDecoder(resp.Body).Decode(&config); err != nil {
		log.Fatal("Error when decoding information about the Configurations")
	}
	go AssignVariables(&config)
	go setUpTmdbLib()
	go getMovieGenresLocally()
	go getShowGenresLocally()
	defer resp.Body.Close()
}

func AssignVariables(config *model.Configuration) {
	variable.ImageUrl = *config.Images.BaseURL
	variable.ImageSize = config.Images.ProfileSizes[len(config.Images.ProfileSizes)-1]
}

func GetTrendingShows(c echo.Context) error {
	var status = http.StatusOK
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	res, err := GetTrending(model.Tv, pageNumber)
	if err != nil {
		log.Println("Error when getting Trending Movies")
		status = http.StatusExpectationFailed
	}
	assignShowGenre(res)
	return c.JSON(status, res)
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
	res, err := GetTrending(model.All, pageNumber)
	if err != nil {
		log.Fatal("Error when getting Trending Movies")
	}
	return c.JSON(http.StatusOK, res)
}

func GetTrending(mediaType string, pageNumber string) (*model.Result, error) {
	url := variable.BaseUrl + "trending/" + mediaType + "/" + model.Day + getApiAuth() + "&page=" + pageNumber
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

func assignShowGenre(result *model.Result) {
	if showsGenres == nil {
		go getShowGenresLocally()
		for showsGenres == nil {
			// Wait for show genres to be downloaded
		}
	}
	for index, movie := range result.Results {
		result.Results[index].Genres = []string{}
		for _, genreId := range movie.GenreIDS {
			result.Results[index].Genres = append(result.Results[index].Genres, getShowGenreFromId(int(genreId)))
		}
		result.Results[index].GenreIDS = nil
	}
}

func getShowGenreFromId(id int) string {
	for _, genre := range showsGenres.Genres {
		if genre.ID == id {
			return genre.Name
		}
	}
	return ""
}

func GetImageSize(c echo.Context) error {
	return c.String(http.StatusOK, variable.ImageSize)
}

func GetImageBaseUrl(c echo.Context) error {
	return c.String(http.StatusOK, variable.ImageUrl)
}

func getApiAuth() string {
	return "?api_key=" + variable.ApiKey
}

func getSearchOptions(c echo.Context) (string, error) {
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	if err != nil {
		log.Println("Error when parsing page number, page number set to default")
	}
	options := fmt.Sprintf("&page=%v&language=%v&region=%v", pageNumber, model.En, "CA")
	return options, err
}

func SearchShows(c echo.Context) error {
	keywords := url.QueryEscape(c.Param(variable.Keywords))
	options, err := getSearchOptions(c)
	if err != nil {
		log.Println("Error during parsing of options, setting variables to default")
	}
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	if err != nil {
		log.Println("Error when getting page number for searching Shows")
	}
	uri := fmt.Sprintf("%vsearch/tv%v%v&query=%v&%s=%v", variable.BaseUrl, getApiAuth(), options, keywords, variable.Page, pageNumber)
	res, err := http.Get(uri)
	var status = http.StatusOK
	var result = &model.Result{}
	if err = json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Println("Error when decoding result from searching for Shows")
		status = http.StatusExpectationFailed
	}
	//res, err := tmdbApi.SearchTv(keywords, options)
	if err != nil {
		log.Println("Error when searching for Shows")
		status = http.StatusExpectationFailed
	}
	assignShowGenre(result)
	return c.JSON(status, result)

}

func GetShowWatchProvider(c echo.Context) error {
	showId := c.Param(variable.Id)
	url := fmt.Sprintf("%vtv/%v/watch/providers%s", variable.BaseUrl, showId, getApiAuth())
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

func GetShowGenres(c echo.Context) error {
	if showsGenres == nil {
		log.Println("Show Genres is nill")
		go getShowGenresLocally()
		return c.JSON(http.StatusExpectationFailed, nil)
	} else {
		return c.JSON(http.StatusOK, showsGenres)
	}
}

func getShowGenresLocally() {
	if tmdbApi != nil {
		options := make(map[string]string)
		options[variable.Language] = model.En
		if res, err := tmdbApi.GetTvGenres(options); err != nil {
			log.Println("Error when getting TV Genres")
			showsGenres = nil
		} else {
			showsGenres = res
		}
	}
}

func DiscoverShows(c echo.Context) error {
	var status = http.StatusOK
	genres := c.QueryParam(variable.Genre)
	providersIds := c.QueryParam(variable.ProvidersIds)
	//keywords := c.QueryParam(variable.Keywords)
	pageNumber, err := getPageNumber(c.QueryParam(variable.Page))
	if err != nil {
		log.Println("Error when parsing page number")
		status = http.StatusExpectationFailed
	}
	uri := fmt.Sprintf("%vdiscover/tv%v&language=%v&watch_region=%v", variable.BaseUrl, getApiAuth(), model.En, model.CA)
	if genres != "" {
		uri = fmt.Sprintf("%v&%v=%v", uri, variable.Genre, genres)
	}
	if providersIds != "" {
		uri = fmt.Sprintf("%v&%v=%v", uri, variable.ProvidersIds, providersIds)
	}
	//if keywords != "" {
	//	uri = fmt.Sprintf("%v&%v=%v", uri, variable.Keywords, keywords)
	//}
	uri = fmt.Sprintf("%v&%v=%v", uri, variable.Page, pageNumber)
	res, err := http.Get(uri)
	if err != nil {
		log.Println("Error during getting discover shows")
		status = http.StatusExpectationFailed
	}
	var results = &model.Result{}

	if err = json.NewDecoder(res.Body).Decode(&results); err != nil {
		log.Println("Error during decoding of discover shows")
		status = http.StatusExpectationFailed
	}
	assignShowGenre(results)
	return c.JSON(status, results)
}

func Test(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}

func GetImageSizes(c echo.Context) error {
	return c.JSON(http.StatusOK, config.Images.LogoSizes)
}

func GetShows(c echo.Context) error {

	showIds := &[]int{}
	if err := c.Bind(&showIds); err != nil {
		log.Println(err.Error())
	}
	if len(*showIds) == 0 {
		return nil
	}
	var shows model.Result
	shows.Results = make([]model.Media, 0)
	for _, id := range *showIds {
		shows.Results = append(shows.Results, getMedia(Movie, id))
	}
	return c.JSON(http.StatusOK, shows)
}

func getMedia(mediaType int, id int) model.Media {
	var uri string
	if mediaType == Movie {
		uri = fmt.Sprintf("%vmovie/%v%v", variable.BaseUrl, id, getApiAuth())
	} else {
		uri = fmt.Sprintf("%stv/%v%v", variable.BaseUrl, id, getApiAuth())
	}
	reqRes, err := http.Get(uri)
	if err != nil {
		log.Println("Error when getting movie detail")
	}
	var result = &model.MediaSearched{}
	if err := json.NewDecoder(reqRes.Body).Decode(result); err != nil {
		log.Println("Error when decoding movie detail")
	}
	return convertToMedia(*result)
}
func convertToMedia(res model.MediaSearched) model.Media {
	var media = model.Media{ReleaseDate: res.ReleaseDate, Adult: res.Adult, BackdropPath: res.BackdropPath, OriginalLanguage: res.OriginalLanguage, PosterPath: res.PosterPath, Title: res.Title, ID: res.ID, Overview: res.Overview, OriginCountry: res.OriginCountry, Name: res.Name, FirstAirDate: res.FirstAirDate}
	for _, genre := range res.GenresSearched {
		media.Genres = append(media.Genres, genre.Name)
	}
	return media
}
