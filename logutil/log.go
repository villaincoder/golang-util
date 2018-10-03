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

var level = INFO

type Config struct {
	Level LogType
}

func LoadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{
			Level: INFO,
		}
	}
	level := util.GetEnvStr("LOG_LEVEL", "")
	if level != "" {
		config.Level = tag2Type(level)
	}
	return config
}

func Init(config *Config) {
	if config == nil {
		config = LoadEnvConfig(config)
	}
	level = config.Level
}

func type2Tag(logType LogType) (logTag string) {
	switch logType {
	case VERBOSE:
		logTag = "VERBOSE"
	case DEBUG:
		logTag = "DEBUG"
	case INFO:
		logTag = "INFO"
	case WARN:
		logTag = "WARN"
	case ERR:
		logTag = "ERR"
	case FATAL:
		logTag = "FATAL"
	default:
		logTag = "INFO"
	}
	return
}

func tag2Type(logTag string) (logType LogType) {
	switch logTag {
	case "VERBOSE":
		logType = VERBOSE
	case "DEBUG":
		logType = DEBUG
	case "INFO":
		logType = INFO
	case "WARN":
		logType = WARN
	case "ERR":
		logType = ERR
	case "FATAL":
		logType = FATAL
	default:
		logType = INFO
	}
	return
}

func getLogPrefix(logType LogType) string {
	logTag := type2Tag(logType)
	return fmt.Sprintf("[%s]:\n%s\n", logTag, util.GetCallerStack(2))
}

func Verbosef(format string, v ...interface{}) {
	if level > VERBOSE {
		return
	}
	log.Printf("%s%s\n", getLogPrefix(VERBOSE), fmt.Sprintf(format, v...))
}

func Debugf(format string, v ...interface{}) {
	if level > DEBUG {
		return
	}
	log.Printf("%s%s\n", getLogPrefix(DEBUG), fmt.Sprintf(format, v...))
}

func Infof(format string, v ...interface{}) {
	if level > INFO {
		return
	}
	log.Printf("%s%s\n", getLogPrefix(INFO), fmt.Sprintf(format, v...))
}

func Warnf(format string, v ...interface{}) {
	if level > WARN {
		return
	}
	log.Printf("%s%s\n", getLogPrefix(WARN), fmt.Sprintf(format, v...))
}

func Warn(err error) {
	if level > WARN {
		return
	}
	log.Printf("%s%+v\n", getLogPrefix(WARN), err)
}

func Errf(format string, v ...interface{}) {
	if level > ERR {
		return
	}
	log.Printf("%s%s\n", getLogPrefix(ERR), fmt.Sprintf(format, v...))
}

func Err(err error) {
	if level > ERR {
		return
	}
	log.Printf("%s%+v\n", getLogPrefix(ERR), err)
}

func Fatalf(format string, v ...interface{}) {
	log.Printf("%s%s\n", getLogPrefix(FATAL), fmt.Sprintf(format, v...))
}

func Fatal(err error) {
	log.Printf("%s%+v\n", getLogPrefix(FATAL), err)
}
