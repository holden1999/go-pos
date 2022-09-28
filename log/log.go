package log

import (
	"github.com/rs/zerolog"
	"os"
)

var Log *zerolog.Logger

func New(loglevel string) {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if loglevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = &logger
}
