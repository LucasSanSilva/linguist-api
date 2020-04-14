package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Lesson struct {
	gorm.Model
	Language string	`gorm:"type:varchar(3)"`
	Title    string	`gorm:"type:varchar(256)"`
	Content  string	`gorm:"type:text"`
	Media    string	`gorm:"type:varchar(256)"`
}

type Lessons interface {
	CreateLesson(c echo.Context, lesson Lesson) error
}

type lessonsRepository struct {
	db *gorm.DB
}

func NewLessonsRepository() Lessons {
	return &lessonsRepository{gormClient}
}

func (l *lessonsRepository) CreateLesson(c echo.Context, lesson Lesson) error {
	query := l.db.Create(&lesson)
	return query.Error
}
