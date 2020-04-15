package service

import (
	"github.com/labstack/echo"
	"github.com/santoslucas/linguist-api/repository"
)

type Lesson struct {
	Language string `json:"language"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Media    string `json:"media"`
}

type Lessons interface {
	CreateLesson(c echo.Context, lesson Lesson) error
	GetLessons(c echo.Context, language string) ([]Lesson, error)
}

type lessonsService struct {
	lessonsRepository	repository.Lessons
}

func NewLessonsService() Lessons {
	return &lessonsService{
		lessonsRepository: repository.NewLessonsRepository(),
	}
}

func (l *lessonsService) CreateLesson(c echo.Context, lesson Lesson) error {
	repoLesson := repository.Lesson{
		Language: lesson.Language,
		Title:    lesson.Title,
		Content:  lesson.Content,
		Media:    lesson.Media,
	}

	err := l.lessonsRepository.CreateLesson(c, repoLesson)

	return err
}

func (l *lessonsService) GetLessons(c echo.Context, language string) ([]Lesson, error) {
	var lessons []Lesson

	lessonsData, err := l.lessonsRepository.GetLessons(c, language)
	if err != nil {
		return lessons, err
	}

	for i := 0; i < len(lessonsData); i++ {
		lesson := Lesson{
			Language:	lessonsData[i].Language,
			Title:	lessonsData[i].Title,
			Content: lessonsData[i].Content,
			Media:	lessonsData[i].Media,
		}
		lessons = append(lessons, lesson)
	}

	return lessons, nil
}
