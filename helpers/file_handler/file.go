package file_handler

import "github.com/maziz/gCore/helpers"

type FileHandler struct{}

const fileHandlerKey helpers.ContextKey = "FileHandlerContext"

type File interface {
	FilePathExists(file string) bool
}
