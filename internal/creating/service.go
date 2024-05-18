package creating

import (
	"context"
	mooc "hexagonal-go-kata/internal"
)

type CourseService struct {
	courseRepository mooc.CourseRepository
}

func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

func (s CourseService) CreateCourse(ctxt context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	return s.courseRepository.Save(ctxt, course)
}
