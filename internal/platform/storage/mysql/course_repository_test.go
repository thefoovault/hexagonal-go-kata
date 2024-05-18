package mysql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	mooc "hexagonal-go-kata/internal"
	"testing"
)

func TestCourseRepository_SaveFailed(t *testing.T) {
	courseId, courseName, courseDuration := "8a1c5cdc-ba57-445a-994d-aa412d23723f", "Demo course", "10 months"
	course, err := mooc.NewCourse(courseId, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseId, courseName, courseDuration).
		WillReturnError(errors.New("Something failed"))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.Error(t, err)
}

func TestCourseRepository_Save(t *testing.T) {
	courseId, courseName, courseDuration := "8a1c5cdc-ba57-445a-994d-aa412d23723f", "Demo course", "10 months"
	course, err := mooc.NewCourse(courseId, courseName, courseDuration)
	require.NoError(t, err)

	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)

	sqlMock.ExpectExec(
		"INSERT INTO courses (id, name, duration) VALUES (?, ?, ?)").
		WithArgs(courseId, courseName, courseDuration).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo := NewCourseRepository(db)

	err = repo.Save(context.Background(), course)

	assert.NoError(t, sqlMock.ExpectationsWereMet())
	assert.NoError(t, err)
}
