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
	Debug(op string, msg string, args ...any)
	Info(op string, msg string, args ...any)
	Warn(op string, msg string, args ...any)
	Error(op string, msg string, args ...any)
	Panic(op string, msg string, args ...any)
	DebugContext(ctx context.Context, op string, msg string, keysAndVals ...any)
	InfoContext(ctx context.Context, op string, msg string, keysAndVals ...any)
	WarnContext(ctx context.Context, op string, msg string, keysAndVals ...any)
	ErrorContext(ctx context.Context, op string, msg string, keysAndVals ...any)
	PanicContext(ctx context.Context, op string, msg string, keysAndVals ...any)
}
