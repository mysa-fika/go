package log

import (
	"context"
	"io"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func NewZerolog(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	return &Logger{logger: &logger}
}

func NewConsole(isDebug bool) *Logger {
	logLevel := zerolog.InfoLevel
	if isDebug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{logger: &logger}
}

func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

func (l *Logger) With() zerolog.Context {
	return l.logger.With()
}

func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return l.logger.Level(level)
}

func (l *Logger) Sample(s zerolog.Sampler) zerolog.Logger {
	return l.logger.Sample(s)
}

func (l *Logger) Hook(h zerolog.Hook) zerolog.Logger {
	return l.logger.Hook(h)
}

func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

func (l *Logger) WithLevel(level zerolog.Level) *zerolog.Event {
	return l.logger.WithLevel(level)
}

func (l *Logger) Log() *zerolog.Event {
	return l.logger.Log()
}

func (l *Logger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

func (l *Logger) Ctx(ctx context.Context) *Logger {
	return &Logger{logger: zerolog.Ctx(ctx)}
}
