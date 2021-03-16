package main

import (
	"fmt"
	"github.com/go-bongo/bongo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/superDeano/account/Dao"
	"github.com/superDeano/account/Handler"
	"log"
)

func main() {
	fmt.Println("test")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	setUpDbConnection()
	setUpRoutes(e)
	e.Logger.Fatal(e.Start(":7000"))
}

func setUpRoutes(e *echo.Echo) {
	e.GET("/addUser", Handler.AddAccount)
	e.GET("/auth", Handler.AuthenticateUser)
	e.GET("/", Handler.Test)
	e.GET("/deleteUser", Handler.DeleteUser)
}

func setUpDbConnection() {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "test",
	}
	var err error
	Dao.Db, err = bongo.Connect(config)

	if err != nil {
		log.Fatal(err)
	}
}
