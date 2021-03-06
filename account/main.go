package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/superDeano/account/Dao"
	"github.com/superDeano/account/Handler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func main() {
	fmt.Println("The Account MicroService")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	setUpDbConnection()
	setUpRoutes(e)
	Handler.ShowsMsBaseUrl = os.Getenv("SHOWS_URL")
	e.Logger.Fatal(e.Start(":8080"))
}

func setUpRoutes(e *echo.Echo) {
	e.PUT("/account/user", Handler.AddAccount)
	e.POST("/account/auth", Handler.AuthenticateUser)
	e.GET("/account/", Handler.Test)
	e.DELETE("/account/user/:"+Dao.AccountUsername, Handler.DeleteUser)
	e.GET("/account/user/:"+Dao.AccountUsername, Handler.GetUser)

	//WatchLists
	e.GET(fmt.Sprintf("/account/user/:%s/watchlist/show", Dao.AccountUsername), Handler.GetShowWatchlist)
	e.GET(fmt.Sprintf("/account/user/:%v/watchlist/movie", Dao.AccountUsername), Handler.GetMovieWatchlist)
	e.PUT(fmt.Sprintf("/account/user/:%v/watchlist/show/:%s", Dao.AccountUsername, Dao.ShowId), Handler.AddShowToWatchlist)
	e.PUT(fmt.Sprintf("/account/user/:%v/watchlist/movie/:%s", Dao.AccountUsername, Dao.MovieId), Handler.AddMovieToWatchlist)
	e.DELETE(fmt.Sprintf("/account/user/:%v/watchlist/show/:%s", Dao.AccountUsername, Dao.ShowId), Handler.RemoveShowFromWatchlist)
	e.DELETE(fmt.Sprintf("/account/user/:%v/watchlist/movie/:%s", Dao.AccountUsername, Dao.MovieId), Handler.RemoveMovieFromWatchlist)
}

func setUpDbConnection() {
	config := options.Client().ApplyURI(os.Getenv("DB_URL"))

	client, err := mongo.Connect(context.TODO(), config)

	if err != nil {
		log.Fatal(err)
	}
	client.Ping(context.TODO(), nil)
	Dao.Db = client.Database(os.Getenv("DB")).Collection(Dao.AccountCollection)
}
