package log_handler

import (
	"context"

	"github.com/Muhammad-Aziz/gCore/utils"
)

const logHandlerKey utils.ContextKey = "LogHandlerContext"

// WithFileChecker adds a FileChecker to the context.
func AddHandlerToContext(ctx context.Context, fc Log) context.Context {
	return context.WithValue(ctx, logHandlerKey, fc)
}

// FileCheckerFromContext retrieves a FileChecker from the context.
func GetHandlerFromContext(ctx context.Context) Log {
	fc, _ := ctx.Value(logHandlerKey).(Log)
	return fc
}

func CheckHandlerIsSetInContext(ctx context.Context) bool {
	_, ok := ctx.Value(logHandlerKey).(Log)
	return ok
}
