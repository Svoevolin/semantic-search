package logger

import (
	"context"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Panic(msg string, args ...any)
	DebugContext(ctx context.Context, msg string, keysAndVals ...any)
	InfoContext(ctx context.Context, msg string, keysAndVals ...any)
	WarnContext(ctx context.Context, msg string, keysAndVals ...any)
	ErrorContext(ctx context.Context, msg string, keysAndVals ...any)
	PanicContext(ctx context.Context, msg string, keysAndVals ...any)
}
