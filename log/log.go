package log

var Log = NewDefaultLogger()

type Level int

const (
	TRACE Level = iota
	DEBUG
	INFO
	WARN
	ERROR
)

type Logger interface {
	SetLevel(level Level)
	Trace(format string, args ...interface{})
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
}

func Trace(format string, args ...interface{}) {
	Log.Trace(format, args...)
}

func Debug(format string, args ...interface{}) {
	Log.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	Log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	Log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	Log.Error(format, args...)
}
