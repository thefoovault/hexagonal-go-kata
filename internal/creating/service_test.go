package creating

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	mooc "hexagonal-go-kata/internal"
	"hexagonal-go-kata/internal/platform/storage/storagemocks"
	"testing"
)

func Test_CourseService_CreateCourse_RepositoryFail(t *testing.T) {
	courseId := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	courseName := "Demo Course"
	courseDuration := "10 months"

	course, err := mooc.NewCourse(courseId, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(errors.New("Unexpected error"))

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseId, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_CourseService_CreateCourse_Succeed(t *testing.T) {
	courseId := "8a1c5cdc-ba57-445a-994d-aa412d23723f"
	courseName := "Demo Course"
	courseDuration := "10 months"

	course, err := mooc.NewCourse(courseId, courseName, courseDuration)
	require.NoError(t, err)

	courseRepositoryMock := new(storagemocks.CourseRepository)
	courseRepositoryMock.On("Save", mock.Anything, course).Return(nil)

	courseService := NewCourseService(courseRepositoryMock)

	err = courseService.CreateCourse(context.Background(), courseId, courseName, courseDuration)

	courseRepositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}
