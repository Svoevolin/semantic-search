package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/delivery/http/handlers"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/delivery/http/middlewares"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
)

type PublicServer struct {
	cfg    *config.App
	echo   *echo.Echo
	logger logger.Logger
}

func NewPublicServer(container *Container) *PublicServer {
	server := &PublicServer{cfg: container.Config, logger: container.Logger}
	return server
}

func (s *PublicServer) Configure(container *Container) (*PublicServer, error) {
	s.echo = echo.New()

	s.echo.Use(
		echoMiddleware.Recover(),
		middlewares.PublicServerCORSMiddleware(container.Config),
	)

	s.v1(container)
	return s, nil
}

func (s *PublicServer) Start() error {
	const op = "app.public_server.Start"

	if s.echo == nil {
		return errors.New(op + ": didn't init echo")
	}

	return s.echo.Start(fmt.Sprintf(":%s", s.cfg.Port))
}

func (s *PublicServer) Shutdown(ctx context.Context) error {
	const op = "app.public_server.Shutdown"

	if s.echo == nil {
		return errors.New(op + ": didn't init echo")
	}
	return s.echo.Shutdown(ctx)
}

func (s *PublicServer) Echo() *echo.Echo {
	return s.echo
}

func (s *PublicServer) v1(c *Container) {
	s.echo.File("/swagger.json", "api/swagger.json")
	// Handler init
	documentHandler := handlers.NewDocument(c.DocumentService, c.Logger)

	// Handler register
	v1 := s.echo.Group("/api/v1")

	v1.POST("/documents", documentHandler.List)
	v1.POST("/documents/upload", documentHandler.Upload)
}
