package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"hexagonal-go-kata/internal/creating"
	"hexagonal-go-kata/internal/platform/bus/inmemory"
	"hexagonal-go-kata/internal/platform/server"
	"hexagonal-go-kata/internal/platform/storage/mysql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	host            = "0.0.0.0"
	port            = 8080
	shutdownTimeout = 10 * time.Second

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

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
