package error_handler

import "github.com/maziz/gCore/helpers"

const errorHandlerKey helpers.ContextKey = "ErrorHandlerContext"

type Error interface {
	HandleError(err error)
}
