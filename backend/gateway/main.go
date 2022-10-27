package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.XML(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/hello", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Hello world</h1>")
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
