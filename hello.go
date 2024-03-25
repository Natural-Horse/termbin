package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":2022"))
}
