package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	var N int
	flag.IntVar(&N, "N", 1, "Workers num")
	flag.Parse()

	server := echo.New()

	server.POST("/addTask", addTask)

	server.Logger.Fatal(server.Start(":8080"))
}

func addTask(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}
