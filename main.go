package main

import (
	"net/http"
	"sample-web/app/user"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/api/user", user.AddUser)
	e.GET("/api/user", user.GetUser)

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
