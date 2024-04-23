package context_handler

import (
	"context"

	"github.com/Muhammad-Aziz/gCore/utils"
)

type DIContext interface {
	GetKey() string
}

type Icontext interface {
	AddHandlerToContext(object DIContext)
	GetHandlerFromContext(key string) DIContext
	CheckHandlerIsSetInContext(key string) bool
	GetContext() context.Context
	AddUtil(utils utils.Utils)
	LoadUtils()
}

type Context struct {
	Context context.Context
	IsError bool
	Utils   []utils.Utils
}

// CreateNewContext creates a new Context object and return its interface Icontext
func CreateContext() Icontext {
	return &Context{Context: context.Background()}
}

func (ctx *Context) AddUtil(utils utils.Utils) {
	ctx.Utils = append(ctx.Utils, utils)
}

func (ctx *Context) LoadUtils() {
	for _, utl := range ctx.Utils {
		utl.Load()
	}
}

// Context.GetContext retrieves the context value
func (ctx *Context) GetContext() context.Context {
	return ctx.Context
}

func (ctx *Context) SetContext(context context.Context) {
	ctx.Context = context
}

// WithFileChecker adds a FileChecker to the context.
func (ctx *Context) AddHandlerToContext(object DIContext) {
	ctx.SetContext(context.WithValue(ctx.GetContext(), object.GetKey(), object))
}

func (ctx *Context) GetValueFromContext(key string) interface{} {
	return ctx.GetContext().Value(key)
}

// FileCheckerFromContext retrieves a FileChecker from the context.
func (ctx *Context) GetHandlerFromContext(key string) DIContext {
	return ctx.GetValueFromContext(key).(DIContext)
}

// CheckHandlerIsSetInContext verify if the value is already set in context or not
func (ctx *Context) CheckHandlerIsSetInContext(key string) bool {
	value := ctx.GetValueFromContext(key)
	return value != nil
}
