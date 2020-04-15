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
	GetLessons(c echo.Context, language string) ([]Lesson, error)
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

func (l *lessonsRepository) GetLessons(c echo.Context, language string) ([]Lesson, error) {
	var result []Lesson
	_ = l.db.Where("language = ?", language).Order("title asc").Find(&result)

	return result, nil
}
