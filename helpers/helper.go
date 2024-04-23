package helpers

import "context"

type ContextKey string

type Helper interface {
	AddHandlerToContext(ctx context.Context) context.Context
	GetHandlerFromContext(ctx context.Context) bool
}
