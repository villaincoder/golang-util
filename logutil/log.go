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
	return fmt.Sprintf("[%s]:\n%s\n", logTag, util.GetCallerStack(3))
}

func Log(logType LogType, err error) {
	if level > logType {
		return
	}
	log.Printf("%s%+v\n", getLogPrefix(logType), err)
}

func Logf(logType LogType, format string, v ...interface{}) {
	if level > logType {
		return
	}
	log.Printf("%s%s\n", getLogPrefix(logType), fmt.Sprintf(format, v...))
}

func Verbosef(format string, v ...interface{}) {
	Logf(VERBOSE, format, v...)
}

func Verbose(err error) {
	Log(VERBOSE, err)
}

func Debugf(format string, v ...interface{}) {
	Logf(DEBUG, format, v...)
}

func Debug(err error) {
	Log(DEBUG, err)
}

func Infof(format string, v ...interface{}) {
	Logf(INFO, format, v...)
}

func Info(err error) {
	Log(INFO, err)
}

func Warnf(format string, v ...interface{}) {
	Logf(WARN, format, v...)
}

func Warn(err error) {
	Log(WARN, err)
}

func Errf(format string, v ...interface{}) {
	Logf(ERR, format, v...)
}

func Err(err error) {
	Log(ERR, err)
}

func Fatalf(format string, v ...interface{}) {
	Logf(FATAL, format, v...)
}

func Fatal(err error) {
	Log(FATAL, err)
}
