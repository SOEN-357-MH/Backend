package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"shows/handler"
	"shows/variable"
)

func main() {
	setUpVariables()
	fmt.Println("The Shows microservice")
	setUpServer()
}

func setUpServer() {
	server := echo.New()
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	setupServerRoutes(server)
	server.Logger.Fatal(server.Start(os.Getenv("SERVER_PORT")))
}

func setupServerRoutes(e *echo.Echo) {
	defer handler.ConfigureSomeKeyVariables()
	baseUrl := "/media/"
	e.GET(baseUrl, handler.Test)
	e.GET(baseUrl+"health", handler.GetHealth)
	e.GET(baseUrl+"baseImageUrl", handler.GetImageBaseUrl)
	e.GET(baseUrl+"imageSize", handler.GetImageSize)
	e.GET(baseUrl+"imageSizes", handler.GetImageSizes)

	e.POST(fmt.Sprintf("%vshows",baseUrl), handler.GetShows)
	e.POST(fmt.Sprintf("%vmovies",baseUrl), handler.GetMovies)

	// Get Trending
	e.GET(baseUrl+"trending/movie/:"+variable.Page, handler.GetTrendingMovies)
	e.GET(baseUrl+"trending/show/:"+variable.Page, handler.GetTrendingShows)
	e.GET(baseUrl+"trending/all/:"+variable.Page, handler.GetTrendingAll)
	e.GET(baseUrl+"trending/person/:"+variable.Page, handler.GetTrendingPersons)

	// Search
	e.GET(fmt.Sprintf("%vsearch/movie/:%v/:%v", baseUrl, variable.Keywords, variable.Page), handler.SearchMovie)
	e.GET(baseUrl+"search/show/:"+variable.Keywords+"/:"+variable.Page, handler.SearchShows)
	e.GET(fmt.Sprintf("%vmovie/genre", baseUrl), handler.GetMovieGenres)
	e.GET(fmt.Sprintf("%vshow/genre", baseUrl), handler.GetShowGenres)

	// Get Watch Providers or streaming platforms
	e.GET(fmt.Sprintf("%vshow/:%v/provider", baseUrl, variable.Id), handler.GetShowWatchProvider)
	e.GET(fmt.Sprintf("%vmovie/:%v/provider", baseUrl, variable.Id), handler.GetMovieWatchProvider)

	// Discover new stuffs based on filters
	//movie/genres/providersIds/keywords/pageNumber
	e.GET(fmt.Sprintf("%vdiscover/movie", baseUrl), handler.DiscoverMovies)
	e.GET(fmt.Sprintf("%vdiscover/show", baseUrl), handler.DiscoverShows)
}

func setUpVariables() {
	variable.ApiKey = os.Getenv("API_KEY")
	variable.BaseUrl = os.Getenv("BASE_URL")
}
