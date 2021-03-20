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
	fmt.Println("The Shows microservice")
	setUpVariables()
	handler.ConfigureSomeKeyVariables()
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
	baseUrl := "/media/"
	e.GET(baseUrl, handler.Test)
	e.GET(baseUrl+"health", handler.GetHealth)
	e.GET(baseUrl+"baseImageUrl", handler.GetImageBaseUrl)
	e.GET(baseUrl+"imageSize", handler.GetImageSize)
	e.GET(baseUrl+"changeKeys", handler.GetChangeKeys)
	e.GET(baseUrl+"trending/movies/:page", handler.GetTrendingMovies)
	e.GET(baseUrl+"trending/tv/:page", handler.GetTrendingShows)
	e.GET(baseUrl+"trending/all/:page", handler.GetTrendingAll)
	e.GET(baseUrl+"trending/person/:page", handler.GetTrendingPersons)

}

func setUpVariables() {
	variable.Api_key = os.Getenv("API_KEY")
	variable.Base_url = os.Getenv("BASE_URL")
}
