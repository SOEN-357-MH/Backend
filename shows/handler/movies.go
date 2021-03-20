package handler

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"shows/model"
	"shows/variable"
)

func SearchMovie(c echo.Context) error {
	var status = http.StatusOK
	keywords := c.Param(variable.Keywords)
	options, err := getSearchOptions(c)
	uri := fmt.Sprintf("%vsearch/movie%v&%v&query=%v", variable.BaseUrl, getApiAuth(), options, keywords)
	resp, err := http.Get(uri)
	if err != nil {
		log.Println("Error when trying to search for movies")
	}
	var results = &model.Result{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		log.Println("Error when parsing results of search for movies")
		status = http.StatusExpectationFailed
	}
	assignMovieGenre(results)
	return c.JSON(status, results)
}

func GetTrendingMovies(c echo.Context) error {
	var status = http.StatusOK
	pageNumber, err := getPageNumber(c.Param(variable.Page))
	res, err := GetTrending(string(model.Movie), pageNumber)
	if err != nil {
		log.Println("Error when getting Trending Movies")
		status = http.StatusInternalServerError
	}
	assignMovieGenre(res)
	return c.JSON(status, res)
}

func GetMovieWatchProvider(c echo.Context) error {
	movieId := c.Param(variable.Id)
	url := fmt.Sprintf("%smovie/%v/watch/providers%s", variable.BaseUrl, movieId, getApiAuth())
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

func GetMovieGenres(c echo.Context) error {
	if movieGenres == nil {
		log.Println("Movie genres is nil")
		go getMovieGenresLocally()
		return c.JSON(http.StatusExpectationFailed, nil)
	}
	return c.JSON(http.StatusOK, movieGenres)
}

func getMovieGenresLocally() {
	options := make(map[string]string)
	options[variable.Language] = string(model.En)
	res, err := tmdbApi.GetMovieGenres(options)
	if err != nil {
		log.Println("Error when getting Movie genres")
		movieGenres = nil
	}
	movieGenres = res
}

func assignMovieGenre(result *model.Result) {
	for index, movie := range result.Results {
		result.Results[index].Genres = []string{}
		for _, genreId := range movie.GenreIDS {
			result.Results[index].Genres = append(result.Results[index].Genres, getMovieGenreFromId(int(genreId)))
		}
		result.Results[index].GenreIDS = nil
	}
}
func getMovieGenreFromId(id int) string {
	for _, genre := range movieGenres.Genres {
		if genre.ID == id {
			return genre.Name
		}
	}
	return ""
}
