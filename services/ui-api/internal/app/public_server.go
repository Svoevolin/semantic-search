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
	echoSwagger "github.com/swaggo/echo-swagger"
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

func (s *PublicServer) Configure(c *Container) (*PublicServer, error) {
	s.echo = echo.New()

	s.echo.Use(
		echoMiddleware.Recover(),
		middlewares.PublicServerCORSMiddleware(c.Config),
		middlewares.RequestIDMiddleware(),
		middlewares.RequestLogger(c.Logger, "Public server request"),
	)

	s.v1(c)
	return s, nil
}

func (s *PublicServer) Start(ctx context.Context) error {
	const op = "app.public_server.Start"

	if s.echo == nil {
		return errors.New(op + ": didn't init echo")
	}

	s.logger.InfoContext(ctx, op, "starting public server", "port", s.cfg.Port)
	return s.echo.Start(fmt.Sprintf(":%s", s.cfg.Port))
}

func (s *PublicServer) Shutdown(ctx context.Context) error {
	const op = "app.public_server.Shutdown"

	if s.echo == nil {
		return errors.New(op + ": didn't init echo")
	}
	s.logger.InfoContext(ctx, op, "shutting down public server")
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

	v1.GET("/swagger/*", echoSwagger.EchoWrapHandler(func(c *echoSwagger.Config) {
		c.URLs = []string{"/swagger.json"}
	}))

	v1.POST("/documents", documentHandler.List)
	v1.POST("/documents/upload", documentHandler.Upload)
}
