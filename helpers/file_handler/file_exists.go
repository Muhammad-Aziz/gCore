package file_handler

import "os"

// FilePathExists Check if a file exists or not
func (fileHandler FileHandler) FilePathExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
