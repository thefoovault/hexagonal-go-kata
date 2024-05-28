package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"hexagonal-go-kata/internal/platform/server/handler/courses"
	"hexagonal-go-kata/internal/platform/server/handler/health"
	"hexagonal-go-kata/kit/command"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	httpAddress     string
	engine          *gin.Engine
	shutdownTimeout time.Duration
	//dependencies
	commandBus command.Bus
}

func New(context context.Context, host string, port uint, shutdownTimeout time.Duration, command command.Bus) (context.Context, Server) {
	srv := Server{
		engine:          gin.New(),
		httpAddress:     fmt.Sprintf("%s:%d", host, port),
		commandBus:      command,
		shutdownTimeout: shutdownTimeout,
	}

	srv.registerRoutes()
	return serverContext(context), srv
}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Server running on %s", s.httpAddress)

	server := &http.Server{
		Addr:    s.httpAddress,
		Handler: s.engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Server shut down", err)
		}
	}()

	<-ctx.Done()
	contextShutdown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return server.Shutdown(contextShutdown)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/course", courses.CreateHandler(s.commandBus))
}

func serverContext(ctx context.Context) context.Context {
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-channel
		cancel()
	}()

	return ctx
}
