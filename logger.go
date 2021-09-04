package envlogdriver

import (
	"context"
	"github.com/rs/zerolog"
	"os"
)

type Logger struct {
	*zerolog.Logger
}

type Event struct {
	*zerolog.Event
}

const LogTraceIDKey = "X-Envlogdriver-Trace-Id"

// NewLogger returns a configured logger for production.
// It outputs info level and above logs with sampling.
func NewLogger() *Logger {

	loglevel := os.Getenv("LOG_LEVEL")
	if loglevel == "ERROR" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if loglevel == "WARN" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if loglevel == "INFO" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{Logger: &logger}
}

// To use method chain we need followings

func (l *Logger) Trace() *Event {
	e := l.Logger.Trace()
	return &Event{e}
}

func (l *Logger) Debug() *Event {
	e := l.Logger.Debug()
	return &Event{e}
}

func (l *Logger) Info() *Event {
	e := l.Logger.Info()
	return &Event{e}
}

func (l *Logger) Warn() *Event {
	e := l.Logger.Warn()
	return &Event{e}
}

func (l *Logger) Error() *Event {
	e := l.Logger.Error()
	return &Event{e}
}

func (l *Logger) Err(err error) *Event {
	e := l.Logger.Error().Err(err)
	return &Event{e}
}

func (l *Logger) Fatal() *Event {
	e := l.Logger.Fatal()
	return &Event{e}
}

func (l *Logger) Panic() *Event {
	e := l.Logger.Panic()
	return &Event{e}
}

func (l *Logger) WithLevel(level zerolog.Level) *Event {
	e := l.Logger.WithLevel(level)
	return &Event{e}
}

func (l *Logger) Log() *Event {
	e := l.Logger.Log()
	return &Event{e}
}

func (l *Logger) Print(v ...interface{}) {
	l.Logger.Print(v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.Logger.Printf(format, v...)
}

func (l Logger) Write(p []byte) (n int, err error) {
	n, err = l.Logger.Write(p)
	return n, err
}

// Ctx adds trace id to log event.
func (e *Event) Ctx(ctx context.Context) *Event {
	if v, ok := ctx.Value(LogTraceIDKey).(string); ok && len(v) > 0 {
		return &Event{e.Str("trace", v)}
	}
	return e
}
