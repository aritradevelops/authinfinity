// TODO: json transport for production
package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var instance = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

func Debug() func() *zerolog.Event {
	return instance.Debug
}

func Info() func() *zerolog.Event {
	return instance.Info
}

func Warn() func() *zerolog.Event {
	return instance.Warn
}

func Error() func() *zerolog.Event {
	return instance.Error
}

func Fatal() func() *zerolog.Event {
	return instance.Fatal
}

func Panic() func() *zerolog.Event {
	return instance.Panic
}
