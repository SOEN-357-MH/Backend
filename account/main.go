package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/superDeano/account/Handler"
)

func main() {
	fmt.Println("test")
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	setUpRoutes(e)
	e.Logger.Fatal(e.Start(":7000"))
}

func setUpRoutes(e *echo.Echo) {
	e.GET("/addUser", Handler.AddAccount)
	e.GET("/auth", Handler.AuthenticateUser)
	e.GET("/", Handler.Test)
	e.GET("/deleteUser", Handler.DeleteUser)
}
