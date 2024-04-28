package mooc

import "context"

// Course represents a course data structure
type Course struct {
	id       string
	name     string
	duration string
}

func NewCourse(id, name, duration string) Course {
	return Course{
		id,
		name,
		duration,
	}
}

// CourseRepository defines the contract for the course storage
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

func (c Course) Id() string {
	return c.id
}

func (c Course) Name() string {
	return c.name
}

func (c Course) Duration() string {
	return c.duration
}
