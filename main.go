package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	server := echo.New()

	server.POST("/addTask", addTask)

	server.Logger.Fatal(server.Start(":8080"))
}

func addTask(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}
