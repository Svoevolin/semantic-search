package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
)

func PublicServerCORSMiddleware(cfg *config.App) echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: cfg.PublicServer.AllowOrigins,
		AllowMethods: cfg.PublicServer.AllowMethods,
	})
}
