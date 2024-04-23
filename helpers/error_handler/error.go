package error_handler

import "github.com/Muhammad-Aziz/gCore/helpers"

const errorHandlerKey helpers.ContextKey = "ErrorHandlerContext"

type Error interface {
	HandleError(err error)
}
