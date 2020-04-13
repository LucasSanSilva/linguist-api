package main

import (
	"github.com/labstack/echo"
	"github.com/santoslucas/linguist-api/handler"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/lessons", handler.Lessons.CreateLesson)

	e.Logger.Fatal(e.Start(":7777"))
}
