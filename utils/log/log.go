package log

import (
	"fmt"
	"strconv"
)

const (
	NONE       = 0
	NORMAL     = 1
	DEBUG      = 2
	DIAGNOSTIC = 3
	NEVER      = 99
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

func SetLogLevelFromArgs(argMap map[string]string) {
	logLevelStr, found := argMap["log"]
	if found {
		logLevel, err := NORMAL, error(nil)

		switch logLevelStr {
		case "NONE", "None", "none":
			logLevel = NONE
		case "NORMAL", "Normal", "normal":
			logLevel = NORMAL
		case "DEBUG", "Debug", "debug":
			logLevel = DEBUG
		case "DIAG", "DIAGNOSTIC", "Diag", "Diagnostic", "diag", "diagnostic":
			logLevel = DIAGNOSTIC
		default:
			logLevel, err = strconv.Atoi(logLevelStr)
		}

		if err != nil {
			panic("Error when parsing log level")
		}
		SetLogLevel(logLevel)
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
