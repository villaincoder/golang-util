package util

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

func ToUint64(value interface{}) (uint64Value uint64, err error) {
	if value == nil {
		return
	}
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		uint64Value = uint64(v.Int())
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		uint64Value = v.Uint()
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		uint64Value = uint64(v.Float())
	case reflect.String:
		uint64Value, err = strconv.ParseUint(reflect.ValueOf(value).String(), 10, 64)
	case reflect.Ptr:
		if !v.IsNil() {
			uint64Value, err = ToUint64(v.Elem().Interface())
		}
	default:
		err = errors.New(fmt.Sprintf("not support value %v %T", value, value))
	}
	return
}

func ToUint32(value interface{}) (uint32Value uint32, err error) {
	uint64Value, err := ToUint64(value)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	uint32Value = uint32(uint64Value)
	return
}

func ToInt64(value interface{}) (int64Value int64, err error) {
	uint64Value, err := ToUint64(value)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	int64Value = int64(uint64Value)
	return
}

func ToInt32(value interface{}) (int32Value int32, err error) {
	uint64Value, err := ToUint64(value)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	int32Value = int32(uint64Value)
	return
}

func ToFloat64(value interface{}) (float64Value float64, err error) {
	if value == nil {
		return
	}
	v := reflect.ValueOf(value)
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		float64Value = float64(v.Int())
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		float64Value = float64(v.Uint())
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		float64Value = v.Float()
	case reflect.String:
		float64Value, err = strconv.ParseFloat(reflect.ValueOf(value).String(), 64)
	case reflect.Ptr:
		if !v.IsNil() {
			float64Value, err = ToFloat64(v.Elem().Interface())
		}
	default:
		err = errors.New(fmt.Sprintf("not support value %v %T", value, value))
		break
	}
	return
}

func ToFloat32(value interface{}) (float32Value float32, err error) {
	float64Value, err := ToFloat64(value)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	float32Value = float32(float64Value)
	return
}

func IntFallback(i, fallback int) int {
	if i == 0 {
		return fallback
	}
	return i
}

func UintFallback(i, fallback uint) uint {
	if i == 0 {
		return fallback
	}
	return i
}

func Int64Fallback(i, fallback int64) int64 {
	if i == 0 {
		return fallback
	}
	return i
}

func Uint64Fallback(i, fallback uint64) uint64 {
	if i == 0 {
		return fallback
	}
	return i
}
