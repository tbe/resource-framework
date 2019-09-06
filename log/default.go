package log

import (
	"fmt"
	"os"

	ct "github.com/daviddengcn/go-colortext"
)

type DefaultLogger struct {
	level Level
}

func NewDefaultLogger() Logger {
	ct.Writer = os.Stderr
	return &DefaultLogger{
		level: INFO,
	}
}

func (d *DefaultLogger) SetLevel(level Level) {
	d.level = level
}

func (d *DefaultLogger) Trace(format string, args ...interface{}) {
	if d.level > TRACE {
		return
	}
	ct.Foreground(ct.Blue, false)
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	ct.ResetColor()
}

func (d *DefaultLogger) Debug(format string, args ...interface{}) {
	if d.level > DEBUG {
		return
	}
	ct.Foreground(ct.Blue, true)
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	ct.ResetColor()
}

func (d *DefaultLogger) Info(format string, args ...interface{}) {
	if d.level > INFO {
		return
	}
	ct.Foreground(ct.Green, false)
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	ct.ResetColor()
}

func (d *DefaultLogger) Warn(format string, args ...interface{}) {
	if d.level > WARN {
		return
	}
	ct.Foreground(ct.Red, false)
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	ct.ResetColor()
}

func (d *DefaultLogger) Error(format string, args ...interface{}) {
	ct.Background(ct.Red, false)
	ct.Foreground(ct.White, false)
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
	ct.ResetColor()
	os.Exit(1)
}
