package error_handler

import (
	"context"
)

// WithErrorHandler adds an ErrorHandler to the context.
func (errHandler *ErrorHandler) AddHanlderToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, errorHandlerKey, errHandler)
}

// ErrorHandlerFromContext retrieves an ErrorHandler from the context.
func (errHandler *ErrorHandler) GetHandlerFromContext(ctx context.Context) bool {
	var ok bool
	*errHandler, ok = ctx.Value(errorHandlerKey).(ErrorHandler)
	return ok
}
