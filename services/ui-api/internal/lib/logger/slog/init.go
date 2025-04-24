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
	opRecord   = "operation"
	LevelPanic = slog.Level(12)
)

type Attribute = slog.Attr

func NewAttribute(key string, value any) Attribute {
	return slog.Any(key, value)
}

func Err(err error) slog.Attr {
	return slog.Any("error", err)
}

type SlogLogger struct {
	log *slog.Logger
}

func NewLogger(cfg *config.App, attributes ...Attribute) *SlogLogger {
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
		base = prettylog.New(
			opts,
			prettylog.WithDestinationWriter(os.Stdout),
			prettylog.WithColor(),
		)
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

	base.WithAttrs(attributes)

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

func (l *SlogLogger) Debug(op string, msg string, keysAndVals ...any) {
	l.log.Debug(msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) Info(op string, msg string, keysAndVals ...any) {
	l.log.Info(msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) Warn(op string, msg string, keysAndVals ...any) {
	l.log.Warn(msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) Error(op string, msg string, keysAndVals ...any) {
	l.log.Error(msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) Panic(op string, msg string, keysAndVals ...any) {
	l.log.Log(context.Background(), LevelPanic, msg, append([]any{opRecord, op}, keysAndVals...)...)
	panic(msg)
}

func (l *SlogLogger) DebugContext(ctx context.Context, op string, msg string, keysAndVals ...any) {
	l.log.DebugContext(ctx, msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) InfoContext(ctx context.Context, op string, msg string, keysAndVals ...any) {
	l.log.InfoContext(ctx, msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) WarnContext(ctx context.Context, op string, msg string, keysAndVals ...any) {
	l.log.WarnContext(ctx, msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) ErrorContext(ctx context.Context, op string, msg string, keysAndVals ...any) {
	l.log.ErrorContext(ctx, msg, append([]any{opRecord, op}, keysAndVals...)...)
}

func (l *SlogLogger) PanicContext(ctx context.Context, op string, msg string, keysAndVals ...any) {
	l.log.Log(ctx, LevelPanic, msg, append([]any{opRecord, op}, keysAndVals...)...)
	panic(msg)
}
