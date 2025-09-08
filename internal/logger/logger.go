package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func InitLogger() {
	// Pretty print console logger
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	Log = zerolog.New(output).With().Timestamp().Logger()
}
