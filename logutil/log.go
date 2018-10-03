package logutil

import (
	"fmt"
	"log"

	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type LogType uint32

const (
	VERBOSE LogType = iota
	DEBUG
	INFO
	WARN
	ERR
	FATAL
)

var Level = INFO

func getLogTag(logType LogType) string {
	tagName := "???"
	switch logType {
	case VERBOSE:
		tagName = "VERBOSE"
	case DEBUG:
		tagName = "DEBUG"
	case INFO:
		tagName = "INFO"
	case WARN:
		tagName = "WARN"
	case ERR:
		tagName = "ERR"
	case FATAL:
		tagName = "FATAL"
	}
	return fmt.Sprintf("[%s]:\n%s\n", tagName, util.GetCallerStack(2))
}

func Verbosef(format string, v ...interface{}) {
	if Level > VERBOSE {
		return
	}
	log.Printf("%s%s\n", getLogTag(VERBOSE), fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	if Level > DEBUG {
		return
	}
	log.Printf("%s%s\n", getLogTag(DEBUG), fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	if Level > INFO {
		return
	}
	log.Printf("%s%s\n", getLogTag(INFO), fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	if Level > WARN {
		return
	}
	log.Printf("%s%s\n", getLogTag(WARN), fmt.Sprintf(format, v...))
}

func Warn(err error) {
	if Level > WARN {
		return
	}
	log.Printf("%s%+v\n", getLogTag(WARN), err)
}

func Errf(format string, v ...interface{}) {
	if Level > ERR {
		return
	}
	log.Printf("%s%s\n", getLogTag(ERR), fmt.Sprintf(format, v...))
}

func Err(err error) {
	if Level > ERR {
		return
	}
	log.Printf("%s%+v\n", getLogTag(ERR), err)
}

func Fatalf(format string, v ...interface{}) {
	log.Printf("%s%s\n", getLogTag(FATAL), fmt.Sprintf(format, v...))
}

func Fatal(err error) {
	log.Printf("%s%+v\n", getLogTag(FATAL), err)
}
