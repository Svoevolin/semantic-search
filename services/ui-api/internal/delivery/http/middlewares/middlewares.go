package middlewares

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger"
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
		c.SetRequest(c.Request().WithContext(ctx))
	}

	return middleware.RequestIDWithConfig(cfg)
}

func RequestLogger(logger logger.Logger, message string) echo.MiddlewareFunc {
	const op = "middlewares.RequestLogger"
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogHost:      true,
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogError:     true,
		LogProtocol:  true,
		LogUserAgent: true,
		LogLatency:   true,
		LogRoutePath: true,
		LogRemoteIP:  true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			ctx := c.Request().Context()

			logData := []any{
				"protocol", v.Protocol,
				"host", v.Host,
				"uri", v.URI,
				"route", v.RoutePath,
				"method", v.Method,
				"remote_ip", v.RemoteIP,
				"user_agent", v.UserAgent,
				"status", v.Status,
				"latency", v.Latency,
			}

			if v.Error != nil {
				logger.ErrorContext(ctx, op, v.Error.Error(), logData...)
				return nil
			}

			logger.InfoContext(ctx, op, message, logData...)

			return nil
		},
	})
}
