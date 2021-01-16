package web

import (
	"os"

	"github.com/labstack/echo"
)

func validateUser(username, password string, c echo.Context) (bool, error) {
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")
	if username == user && password == pass {
		return true, nil
	}
	return false, nil
}
