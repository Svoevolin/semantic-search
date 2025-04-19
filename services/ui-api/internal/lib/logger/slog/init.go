package sl

import (
	"context"
	"log/slog"
	"os"
	"strings"

	"github.com/svoevolin/semantic-search/services/ui-api/internal/config"
	"github.com/svoevolin/semantic-search/services/ui-api/internal/lib/logger/slog/handler"
	"github.com/sytallax/prettylog"
)

const (
	LevelPanic = slog.Level(12)
)

type SlogLogger struct {
	log *slog.Logger
}

func NewLogger(cfg *config.App) *SlogLogger {
	level := setupLevel(cfg.Level)

	// Логи выбрасываются, их не нужно мокать (для тестов)
	if cfg.Discard {
		logger := slog.New(handler.NewDiscardHandler())
		slog.SetDefault(logger)
		return &SlogLogger{log: logger}
	}

	var base slog.Handler

	opts := &slog.HandlerOptions{
		Level:       level,
		AddSource:   false,
		ReplaceAttr: nil,
	}

	if cfg.Pretty {
		base = prettylog.New(opts)
	} else {
		switch strings.ToLower(cfg.Format) {
		case "json":
			base = slog.NewJSONHandler(os.Stdout, opts)
		case "text":
			base = slog.NewTextHandler(os.Stdout, opts)
		default:
			base = slog.NewJSONHandler(os.Stdout, opts)
		}
	}
	// Оборачиваем в CtxHandler для поддержки контекстных данных
	logger := slog.New(handler.NewCtxHandler(base))
	slog.SetDefault(logger)
	return &SlogLogger{log: logger}
}

func setupLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *SlogLogger) Debug(ctx context.Context, msg string, keysAndVals ...any) {
	l.log.DebugContext(ctx, msg, keysAndVals...)
}

func (l *SlogLogger) Info(ctx context.Context, msg string, keysAndVals ...any) {
	l.log.InfoContext(ctx, msg, keysAndVals...)
}

func (l *SlogLogger) Warn(ctx context.Context, msg string, keysAndVals ...any) {
	l.log.WarnContext(ctx, msg, keysAndVals...)
}

func (l *SlogLogger) Error(ctx context.Context, msg string, keysAndVals ...any) {
	l.log.ErrorContext(ctx, msg, keysAndVals...)
}

func (l *SlogLogger) Panic(ctx context.Context, msg string, keysAndVals ...any) {
	l.log.Log(ctx, LevelPanic, msg, keysAndVals...)
	panic(msg)
}
