package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hexagonal-go-kata/internal/platform/server/handler/courses"
	"hexagonal-go-kata/internal/platform/server/handler/health"
	"hexagonal-go-kata/kit/command"
	"log"
)

type Server struct {
	httpAddress string
	engine      *gin.Engine
	//dependencies
	commandBus command.Bus
}

func New(host string, port uint, command command.Bus) Server {
	srv := Server{
		engine:      gin.New(),
		httpAddress: fmt.Sprintf("%s:%d", host, port),
		commandBus:  command,
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
	s.engine.POST("/course", courses.CreateHandler(s.commandBus))
}
