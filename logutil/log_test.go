package logutil

import (
	"testing"

	"github.com/pkg/errors"
)

func TestVerbosef(t *testing.T) {
	level = FATAL
	Verbosef("verbose 1")
	level = VERBOSE
	Verbosef("verbose 2")
}

func TestDebugf(t *testing.T) {
	level = FATAL
	Debugf("debug 1")
	level = DEBUG
	Debugf("debug 2")
}

func TestInfof(t *testing.T) {
	level = FATAL
	Infof("info 1")
	level = INFO
	Infof("info 2")
}

func TestWarnf(t *testing.T) {
	level = FATAL
	Warnf("warn 1")
	level = WARN
	Warnf("warn 2")
}

func TestWarn(t *testing.T) {
	level = FATAL
	Warn(errors.Wrap(errors.New("warn1"), "wrap"))
	level = WARN
	Warn(errors.Wrap(errors.New("warn2"), "wrap"))
}

func TestErrf(t *testing.T) {
	level = FATAL
	Errf("err 1")
	level = ERR
	Errf("err 2")
}

func TestErr(t *testing.T) {
	level = FATAL
	Err(errors.Wrap(errors.New("err1"), "wrap"))
	level = ERR
	Err(errors.Wrap(errors.New("err2"), "wrap"))
}
