package main

import (
	"json-bin/controllers"
	"json-bin/firebase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Init firebase
	firebase.InitializeApp()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/", controllers.NewRecord)
	// e.GET("/", newRecord)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
