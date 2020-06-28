package user

import (
	"net/http"

	"github.com/labstack/echo"
)

// AddUser controller method for adding a new user
func AddUser(c echo.Context) error {
	user := new(User)

	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// GetUser controller method for getting a user
func GetUser(c echo.Context) error {
	user := &User{
		"Favour Harrison",
		"hyper_debugger",
		"hyper_debugger@hey.com",
		"password",
	}

	return c.JSONPretty(http.StatusOK, user, "  ")
}
