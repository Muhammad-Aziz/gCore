package file_handler

import "github.com/Muhammad-Aziz/gCore/helpers"

type FileHandler struct{}

const fileHandlerKey helpers.ContextKey = "FileHandlerContext"

type File interface {
	FilePathExists(file string) bool
}
