package mooc

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var ErrInvalidCourseId = errors.New("invalid Course Id")

type CourseId struct {
	value string
}

func NewCourseId(value string) (CourseId, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return CourseId{}, fmt.Errorf("%s, %s", ErrInvalidCourseId, value)
	}
	return CourseId{
		value: id.String(),
	}, nil
}

func (id CourseId) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name cannot be empty")

type CourseName struct {
	value string
}

func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}
	return CourseName{
		value: value,
	}, nil
}

func (name CourseName) String() string {
	return name.value
}

var ErrEmptyCourseDuration = errors.New("the field Course Duration cannot be empty")

type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyCourseDuration
	}
	return CourseDuration{
		value: value,
	}, nil
}

func (duration CourseDuration) String() string {
	return duration.value
}

// Course represents a course data structure
type Course struct {
	id       CourseId
	name     CourseName
	duration CourseDuration
}

func NewCourse(id, name, duration string) (Course, error) {
	courseId, err := NewCourseId(id)
	if err != nil {
		return Course{}, err
	}

	courseName, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	courseDuration, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       courseId,
		name:     courseName,
		duration: courseDuration,
	}, nil
}

// CourseRepository defines the contract for the course storage
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

func (c Course) Id() CourseId {
	return c.id
}

func (c Course) Name() CourseName {
	return c.name
}

func (c Course) Duration() CourseDuration {
	return c.duration
}
