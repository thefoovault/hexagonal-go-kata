package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/huandu/go-sqlbuilder"
	mooc "hexagonal-go-kata/internal"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

// Save implements the mooc.CourseRepository interface
func (r *CourseRepository) Save(context context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		Id:       course.Id().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	_, err := r.db.ExecContext(context, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
