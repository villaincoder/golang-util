package util

import "runtime"

func GetCallerName(skip int) (name string) {
	name = "???"
	var pc uintptr
	var ok bool
	if pc, _, _, ok = runtime.Caller(skip + 1); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}
