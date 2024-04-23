package file_handler

import (
	"context"
)

// AddHandlerToContext adds a FileHandler to the context.
func (fileHandler *FileHandler) AddHandlerToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, fileHandlerKey, fileHandler)
}

// GetHandlerFromContext retrieves a FileHandler from the context.
func (fileHandler *FileHandler) GetHandlerFromContext(ctx context.Context) bool {
	var ok bool
	*fileHandler, ok = ctx.Value(fileHandlerKey).(FileHandler)
	return ok
}
