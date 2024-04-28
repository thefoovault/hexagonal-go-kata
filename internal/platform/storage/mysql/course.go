package mysql

const (
	sqlCourseTable = "courses"
)

type sqlCourse struct {
	Id       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}
