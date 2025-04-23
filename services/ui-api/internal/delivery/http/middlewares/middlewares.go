package middlewares

import (
	"context"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	slogHandler "github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog/handler"
)

type ContextKeyHeaderReqID string

const HeaderRequestID ContextKeyHeaderReqID = "X-Request-Id"

func PublicServerCORSMiddleware(cfg *config.App) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.AllowOrigins,
		AllowMethods: cfg.AllowMethods,
	})
}

func RequestIDMiddleware() echo.MiddlewareFunc {
	cfg := middleware.DefaultRequestIDConfig

	cfg.Generator = func() string {
		return uuid.New().String()
	}

	cfg.RequestIDHandler = func(c echo.Context, reqID string) {
		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, HeaderRequestID, reqID)             // Для клиентов
		ctx = context.WithValue(ctx, slogHandler.RequestIDLogKey, reqID) // Для логов
	}

	return middleware.RequestIDWithConfig(cfg)
}
