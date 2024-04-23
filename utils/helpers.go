package utils

import "os"

var ShouldExitOnError = true

func SetShouldExitOnError(newShouldExitOnError bool) {
	ShouldExitOnError = newShouldExitOnError
}

func Must(err error) bool {
	isNilError := err != nil
	if isNilError {
		LogError().Err(err).Msg("Fatal error")
		if ShouldExitOnError {
			os.Exit(1)
		}
	}
	return isNilError
}
