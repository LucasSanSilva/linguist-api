package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type Lesson struct {
	Language string `json:"language"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Media    string `json:"media"`
}

type lessonsHandler struct {
}

var Lessons *lessonsHandler

func init() {
	Lessons = &lessonsHandler{}
}

func (l *lessonsHandler) CreateLesson(c echo.Context) error {
	lesson := new(Lesson)
	err := c.Bind(lesson)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	return c.String(http.StatusCreated, "MAOE")
}
