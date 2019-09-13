package logger

import (
	"github.com/jackc/pgx"
	"github.com/rs/zerolog"
)

// Logger is a logger interface for pgx and goa
type Logger struct {
	zerolog.Logger
}

// LogLevel is log severity level
type LogLevel int

// Log is a function to satisfy pgx logger interface
func (l Logger) Log(loglevel pgx.LogLevel, msg string, data map[string]interface{}) {
	var ev *zerolog.Event
	switch loglevel {
	case pgx.LogLevelDebug:
		ev = l.Debug()
	case pgx.LogLevelError:
		ev = l.Error()
	case pgx.LogLevelInfo:
		ev = l.Info()
	case pgx.LogLevelWarn:
		ev = l.Warn()
	default:
		ev = l.Debug()
	}
	for k, v := range data {
		if k == "time" {
			continue
		}
		ev = ev.Interface(k, v)
	}
	ev.Msg(msg)
}

// GoaLogger returns new GoaLogger instance
func (l *Logger) GoaLogger() GoaLogger {
	return GoaLogger{l}
}

// GoaLogger is a wrapper for Logger to satisfy goa LogAdapter interface
type GoaLogger struct {
	logger *Logger
}

// Error sends an Error log message
func (gl GoaLogger) Error(msg string, data ...interface{}) {
	ev := gl.logger.Error()
	gl.log(ev, msg, data...)
}

// Info sends an Info log message
func (gl GoaLogger) Info(msg string, data ...interface{}) {
	ev := gl.logger.Info()
	gl.log(ev, msg, data...)
}

func (gl GoaLogger) log(ev *zerolog.Event, msg string, data ...interface{}) {
	isKey := true
	var key string
	for _, v := range data {
		if isKey {
			var ok bool
			if key, ok = v.(string); !ok {
				break
			}
		} else if key != "" {
			if key == "time" {
				key = "timer"
			}
			ev = ev.Interface(key, v)
		}
		isKey = !isKey
	}
	ev.Msg(msg)
}
