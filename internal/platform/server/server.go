package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_test/internal/platform/server/handler/courses"
	"go_test/internal/platform/server/handler/health"
	"log"
)

type Server struct {
	httpAddress string
	engine      *gin.Engine
}

func New(host string, port uint) Server {
	srv := Server{
		engine:      gin.New(),
		httpAddress: fmt.Sprintf("%s:%d", host, port),
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
	s.engine.POST("/course", courses.CreateHandler())
}
