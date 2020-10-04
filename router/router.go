package router

import (
	"net/http"

	"github.com/labstack/echo"
)

func Init() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
