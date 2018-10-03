package util

import (
	"fmt"
	"runtime"
)

func GetCallerName(skip int) (name string) {
	name = "???"
	if pc, _, _, ok := runtime.Caller(skip + 1); ok {
		name = runtime.FuncForPC(pc).Name()
	}
	return
}

func GetCallerStack(skip int) (stack string) {
	stack = "???"
	if pc, file, line, ok := runtime.Caller(skip + 1); ok {
		stack = fmt.Sprintf("%s\n\t%s:%d", runtime.FuncForPC(pc).Name(), file, line)
	}
	return
}
