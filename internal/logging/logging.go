package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func NewLogger(component string, env string) zerolog.Logger {
	var logger zerolog.Logger
	if env == "dev" {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().
			Timestamp().
			Str("component", component).
			Logger()
	} else {
		logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Str("component", component).
			Logger()
	}
	return logger
}
