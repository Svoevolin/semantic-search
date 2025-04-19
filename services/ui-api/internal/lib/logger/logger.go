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
	Debug(ctx context.Context, msg string, keysAndVals ...any)
	Info(ctx context.Context, msg string, keysAndVals ...any)
	Warn(ctx context.Context, msg string, keysAndVals ...any)
	Error(ctx context.Context, msg string, keysAndVals ...any)
	Panic(ctx context.Context, msg string, keysAndVals ...any)
}
