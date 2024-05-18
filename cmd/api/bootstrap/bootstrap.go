package bootstrap

import (
	"database/sql"
	"fmt"
	"hexagonal-go-kata/internal/creating"
	"hexagonal-go-kata/internal/platform/server"
	"hexagonal-go-kata/internal/platform/storage/mysql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "0.0.0.0"
	port = 8080

	dbUser = "hexagonalGoKata"
	dbPass = "hexagonalGoKata"
	dbHost = "hexagonal-go-kata-mysql"
	dbPort = "3306"
	dbName = "hexagonalGoKata"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	srv := server.New(host, port, creatingCourseService)
	return srv.Run()
}
