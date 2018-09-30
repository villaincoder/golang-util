package util

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type EnumCheckOption struct {
	MaxValue interface{}
	Value    interface{}
}

func getCheckCallerName() (name string) {
	name = "???"
	var pc uintptr
	var ok bool
	if pc, _, _, ok = runtime.Caller(2); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

func CheckArgs(fields map[string]interface{}) (err error) {
	for key, value := range fields {
		switch value.(type) {
		case string:
			if value.(string) == "" {
				err = errors.New(fmt.Sprintf("%s %s is empty string", getCheckCallerName(), key))
			}
			break
		case time.Time:
			if IsInvalidTime(value.(time.Time)) {
				err = errors.New(fmt.Sprintf("%s %s is invalid time", getCheckCallerName(), key))
			}
			break
		default:
			if value == nil || reflect.ValueOf(value).IsNil() {
				err = errors.New(fmt.Sprintf("%s %s is nil", getCheckCallerName(), key))
				return
			}
			break
		}
	}
	return
}

func CheckStrArrayArg(field string, arr []string) (err error) {
	if arr == nil {
		err = errors.New(fmt.Sprintf("%s %s is nil", getCheckCallerName(), field))
		return
	}
	for index, value := range arr {
		if value == "" {
			err = errors.New(fmt.Sprintf("%s %s item %d is empty", getCheckCallerName(), field, index))
			return
		}
	}
	return
}
