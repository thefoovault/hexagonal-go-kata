package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexagonal-go-kata/internal/creating"
	"hexagonal-go-kata/internal/platform/server/handler/courses"
	"hexagonal-go-kata/internal/platform/server/handler/health"
	"log"
)

type Server struct {
	httpAddress           string
	engine                *gin.Engine
	creatingCourseService creating.CourseService
}

func New(host string, port uint, creatingCourseService creating.CourseService) Server {
	srv := Server{
		engine:                gin.New(),
		httpAddress:           fmt.Sprintf("%s:%d", host, port),
		creatingCourseService: creatingCourseService,
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
	s.engine.POST("/course", courses.CreateHandler(s.creatingCourseService))
}
