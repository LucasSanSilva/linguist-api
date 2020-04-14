package handler

import (
	"github.com/labstack/echo"
	"github.com/santoslucas/linguist-api/service"
	"net/http"
)

type lessonsHandler struct {
	lessonsService	service.Lessons
}

var Lessons *lessonsHandler

func init() {
	Lessons = &lessonsHandler{
		lessonsService: service.NewLessonsService(),
	}
}

func (l *lessonsHandler) CreateLesson(c echo.Context) error {
	lesson := new(service.Lesson)
	err := c.Bind(lesson)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}
	
	err = l.lessonsService.CreateLesson(c, *lesson)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusCreated, "MAOE")
}
