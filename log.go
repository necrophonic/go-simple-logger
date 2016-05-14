// Package log defines a generic log implemetation with differing granularity levels.
//
// The logger is built on top of the standard golang log package.
package log

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Level setting constants
//   LevelTrace > LevelDebug > LevelWarn > LevelInfo > LevelError > LevelNone
const (
	LevelNone  = 0
	LevelError = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelDebug = 4
	LevelTrace = 5
)

//Logger is the default interface all logs must implement in this library
type Logger interface {
	Fatal(format string, v ...interface{})
	Trace(format string, v ...interface{})
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}

var (
	level = LevelInfo
)

// DefaultLogger defines a simple logging package
type DefaultLogger struct {
	Level int
}

// Init initialises the logger. This must be called at least once before and of
// the functions can be called.
func Init(l int) error {
	if l < LevelNone || l > LevelTrace {
		return errors.New("unable to init logger: invalid level")
	}
	level = l
	return nil
}

// InitFromString initialises the logger with Init(). Takes a string rather than a numerical levvel
func InitFromString(l string) error {
	switch strings.ToLower(l) {
	case "error":
		Init(LevelError)
	case "info":
		Init(LevelInfo)
	case "warn":
		Init(LevelWarn)
	case "debug":
		Init(LevelDebug)
	case "trace":
		Init(LevelTrace)
	case "none":
		Init(LevelNone)
	default:
		return fmt.Errorf("unable to init logger: unknown level '%s'", l)
	}
	return nil
}

// Fatal provides fatal level logging. Being fatal it will log, and then it will
// exit the current process.
func Fatal(format string, v ...interface{}) {
	if v != nil {
		log.Fatalf("[FATAL] "+format, v...)
	} else {
		log.Fatalln("[FATAL] " + format)
	}
}

// Error provides error level logging
func Error(format string, v ...interface{}) {
	writeLog("[ERROR]", LevelError, format, v...)
}

// Debug provides debug level logging
func Debug(format string, v ...interface{}) {
	writeLog("[DEBUG]", LevelDebug, format, v...)
}

// Warn provides warning level logging
func Warn(format string, v ...interface{}) {
	writeLog("[WARN]", LevelWarn, format, v...)
}

// Info provides info level logging
func Info(format string, v ...interface{}) {
	writeLog("[INFO]", LevelInfo, format, v...)
}

// Trace provides trace level logging
func Trace(format string, v ...interface{}) {
	writeLog("[TRACE]", LevelTrace, format, v...)
}

func writeLog(prefix string, l int, format string, v ...interface{}) {
	if level >= l {
		if v != nil {
			log.Printf(prefix+" "+format, v...)
		} else {
			log.Println(prefix + " " + format)
		}
	}
}
