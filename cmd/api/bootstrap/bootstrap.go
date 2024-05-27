package bootstrap

import (
	"database/sql"
	"fmt"
	"hexagonal-go-kata/internal/creating"
	"hexagonal-go-kata/internal/platform/bus/inmemory"
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

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
