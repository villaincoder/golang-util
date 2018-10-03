package logutil

import (
	"fmt"
	"log"
)

type LogLevel uint32

const (
	VERBOSE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERR
	DISABLE
)

var Level = INFO

func Verbosef(format string, v ...interface{}) {
	if Level > VERBOSE {
		return
	}
	log.Printf("[VERBOSE]:%s\n", fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	if Level > DEBUG {
		return
	}
	log.Printf("[DEBUG]:%s\n", fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	if Level > INFO {
		return
	}
	log.Printf("[LOG]:%s\n", fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	if Level > WARN {
		return
	}
	log.Printf("[WARN]:%s\n", fmt.Sprintf(format, v...))
}

func Warn(err error) {
	if Level > WARN {
		return
	}
	log.Printf("[WARN]:\n%+v\n", err)
}

func Errf(format string, v ...interface{}) {
	if Level > ERR {
		return
	}
	log.Printf("[ERR]:%s\n", fmt.Sprintf(format, v...))
}

func Err(err error) {
	if Level > ERR {
		return
	}
	log.Printf("[ERR]:\n%+v\n", err)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf("[FATAL]:%s\n", fmt.Sprintf(format, v))
}

func Fatal(err error) {
	log.Fatalf("[FATAL]:\n%+v\n", err)
}
