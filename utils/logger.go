package utils

import (
	"os"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger zerolog.Logger

func LoadLogger() {
	logFile := &lumberjack.Logger{
		Filename:   "app.log", // Log file name
		MaxSize:    5,         // Maximum log file size in megabytes
		MaxBackups: 3,         // Maximum number of old log files to retain
		MaxAge:     7,         // Maximum number of days to retain old log files
		Compress:   true,      // Whether to compress old log files
	}

	// todo: close log file after server off
	//defer logFile.Close()

	multi := zerolog.MultiLevelWriter(zerolog.ConsoleWriter{Out: os.Stdout}, logFile)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger = zerolog.New(multi).With().Timestamp().Logger()

	LogInfo().Msg("Logger initiated successfully")
}

func LogInfo() *zerolog.Event {
	return logger.Info()
}

func LogDebug() *zerolog.Event {
	return logger.Debug()
}

func LogWarn() *zerolog.Event {
	return logger.Warn()
}

func LogError() *zerolog.Event {
	return logger.Error()
}
