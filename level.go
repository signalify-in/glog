package glog

type Level string

const (
	Panic   Level = "PANIC:"
	Fatal   Level = "FATAL:"
	Error   Level = "ERROR:"
	Warning Level = "WARNING:"
	Info    Level = "INFO:"
	Debug   Level = "DEBUG:"
	Trace   Level = "TRACE:"
)

const (
	LevelPanic uint8 = iota + 1
	LevelFatal
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
	LevelTrace
)
