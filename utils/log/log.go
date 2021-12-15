package log

import "fmt"

const (
	NONE       = 0
	NORMAL     = 1
	DEBUG      = 2
	DIAGNOSTIC = 3
)

const (
	CONSOLE = 1
	FILE    = 2
)

var LOG_LEVEL = NORMAL
var LOG_TIMESTAMP = false
var LOG_OUTPUT = CONSOLE

func SetLogLevel(level int) {
	if level >= NONE && level <= DIAGNOSTIC {
		LOG_LEVEL = level
	}
}

func Print(minLevel int, a ...interface{}) (int, error) {
	if LOG_LEVEL >= minLevel {
		return fmt.Print(a...)
	}
	return 0, nil
}

func Println(minLevel int, a ...interface{}) (int, error) {
	if LOG_LEVEL >= minLevel {
		return fmt.Println(a...)
	}
	return 0, nil
}

func Printf(minLevel int, format string, a ...interface{}) (int, error) {
	if LOG_LEVEL >= minLevel {
		return fmt.Printf(format, a...)
	}
	return 0, nil
}
