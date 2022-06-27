package logger

import (
	"fmt"
	"io"
	"log"
)

const (
	Error = iota
	Trace
	Debug
)

type Logger struct {
	level byte

	dlogger, tlogger, elogger *log.Logger
}

func (l *Logger) Debug(message string, values ...any) {
	if l.level == Debug {
		l.dlogger.Printf(message, values...)
	}
}

func (l *Logger) Trace(message string, values ...any) {
	if l.level >= Trace {
		l.tlogger.Printf(message, values...)
	}
}

func (l *Logger) Error(message string, values ...any) {
	l.elogger.Printf(message, values...)
}

func New(level byte, out, err io.Writer, component string) *Logger {
	var (
		dlogger = logger(out, "DEBUG", component)
		tlogger = logger(out, "TRACE", component)
		elogger = logger(err, "ERROR", component)
	)
	logger := Logger{
		dlogger: dlogger,
		tlogger: tlogger,
		elogger: elogger,
		level:   level,
	}
	return &logger
}

func logger(out io.Writer, level, component string) *log.Logger {
	lflags := log.Ldate | log.Ltime | log.Lmsgprefix
	p := prefix(level, component)
	return log.New(out, p, lflags)
}

func prefix(level, component string) string {
	prefix := fmt.Sprintf("[ %-8s ] [ %-15s ] ", level, component)
	return prefix
}
