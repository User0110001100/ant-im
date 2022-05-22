package logger

const (
	Trace = 0
	Debug
	Info
	Error
)

var level = Trace

func Level() int {
	return level
}
