package main

import (
	"github.com/labstack/echo"
	"github.com/santoslucas/linguist-api/handler"
	"github.com/santoslucas/linguist-api/repository"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/lessons", handler.Lessons.CreateLesson)
	e.GET("/lessons", handler.Lessons.GetLessons)

	defer repository.CloseConnection()

	e.Logger.Fatal(e.Start(":7777"))
}
