package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	mooc "hexagonal-go-kata/internal"
	"hexagonal-go-kata/internal/platform/server/handler/courses"
	"hexagonal-go-kata/internal/platform/server/handler/health"
	"log"
)

type Server struct {
	httpAddress      string
	engine           *gin.Engine
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:           gin.New(),
		httpAddress:      fmt.Sprintf("%s:%d", host, port),
		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Printf("Server running on %s", s.httpAddress)
	return s.engine.Run(s.httpAddress)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/course", courses.CreateHandler(s.courseRepository))
}
