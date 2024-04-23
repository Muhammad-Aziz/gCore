package error_handler

import "log"

type ErrorHandler struct{}

func (ErrorHandler) HandleError(err error) {
	log.Println(err)
}
