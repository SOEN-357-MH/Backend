package main

import (
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/superDeano/account/Dao"
	"github.com/superDeano/account/Handler"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Account MicroService")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	setUpDbConnection()
	setUpRoutes(e)
	e.Logger.Fatal(e.Start(":7000"))
}

func setUpRoutes(e *echo.Echo) {
	e.PUT("/account/addUser", Handler.AddAccount)
	e.POST("/account/auth", Handler.AuthenticateUser)
	e.GET("/account/", Handler.Test)
	e.DELETE("/account/deleteUser/:"+Dao.AccountUsername, Handler.DeleteUser)
}

func setUpDbConnection() {
	config := &bongo.Config{
		ConnectionString: os.Getenv("DB_URL"),
		Database:         os.Getenv("DB"),
	}
	var err error
	Dao.Db, err = bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}
}
