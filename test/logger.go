package test

import (
	"fmt"

	"github.com/tbe/resource-framework/log"
)

type Logger struct {
	messages map[string][]string
}

func NewLogger() *Logger {
	return &Logger{messages: make(map[string][]string)}
}

func (l *Logger) SetLevel(level log.Level) {
	// useless for our case
}


// Reset clears the internal state of the logger
func (l *Logger) Reset() {
	l.messages = make(map[string][]string)
}

func (l *Logger) Trace(format string, args ...interface{}) {
	l.messages["trace"] = append(l.messages["trace"], fmt.Sprintf(format, args...))
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.messages["debug"] = append(l.messages["debug"], fmt.Sprintf(format, args...))
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.messages["info"] = append(l.messages["info"], fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.messages["warn"] = append(l.messages["warn"], fmt.Sprintf(format, args...))
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.messages["error"] = append(l.messages["error"], fmt.Sprintf(format, args...))
}


