package util

import (
	"strconv"
	"testing"
)

func TestToUint64(t *testing.T) {
	i, i8, i16, i32, i64 := int(1), int8(2), int16(3), int32(4), int64(5)
	ui, ui8, ui16, ui32, ui64 := uint(1), uint8(2), uint16(3), uint32(4), uint64(5)
	f32, f64 := float32(1.1), float64(2.2)
	s1, s2 := "1", "hello"
	struct1 := struct{}{}
	var nil1 *struct{}
	if r, _ := ToUint64(i); r != uint64(i) {
		t.Fatal("int ToUint64 error")
	}
	if r, _ := ToUint64(&i); r != uint64(i) {
		t.Fatal("*int ToUint64 error")
	}
	if r, _ := ToUint64(i8); r != uint64(i8) {
		t.Fatal("int8 ToUint64 error")
	}
	if r, _ := ToUint64(&i8); r != uint64(i8) {
		t.Fatal("*int8 ToUint64 error")
	}
	if r, _ := ToUint64(i16); r != uint64(i16) {
		t.Fatal("int16 ToUint64 error")
	}
	if r, _ := ToUint64(&i16); r != uint64(i16) {
		t.Fatal("*int16 ToUint64 error")
	}
	if r, _ := ToUint64(i32); r != uint64(i32) {
		t.Fatal("int32 ToUint64 error")
	}
	if r, _ := ToUint64(&i32); r != uint64(i32) {
		t.Fatal("*int32 ToUint64 error")
	}
	if r, _ := ToUint64(i64); r != uint64(i64) {
		t.Fatal("int64 ToUint64 error")
	}
	if r, _ := ToUint64(&i64); r != uint64(i64) {
		t.Fatal("*int64 ToUint64 error")
	}
	if r, _ := ToUint64(ui); r != uint64(ui) {
		t.Fatal("uint ToUint64 error")
	}
	if r, _ := ToUint64(ui); r != uint64(ui) {
		t.Fatal("*uint ToUint64 error")
	}
	if r, _ := ToUint64(ui8); r != uint64(ui8) {
		t.Fatal("uint8 ToUint64 error")
	}
	if r, _ := ToUint64(ui8); r != uint64(ui8) {
		t.Fatal("*uint8 ToUint64 error")
	}
	if r, _ := ToUint64(ui16); r != uint64(ui16) {
		t.Fatal("uint16 ToUint64 error")
	}
	if r, _ := ToUint64(ui16); r != uint64(ui16) {
		t.Fatal("*uint16 ToUint64 error")
	}
	if r, _ := ToUint64(ui32); r != uint64(ui32) {
		t.Fatal("uint32 ToUint64 error")
	}
	if r, _ := ToUint64(ui32); r != uint64(ui32) {
		t.Fatal("*uint32 ToUint64 error")
	}
	if r, _ := ToUint64(ui64); r != uint64(ui64) {
		t.Fatal("uint64 ToUint64 error")
	}
	if r, _ := ToUint64(ui64); r != uint64(ui64) {
		t.Fatal("*uint64 ToUint64 error")
	}
	if r, _ := ToUint64(f32); r != uint64(f32) {
		t.Fatal("float32 ToUint64 error")
	}
	if r, _ := ToUint64(f32); r != uint64(f32) {
		t.Fatal("*float32 ToUint64 error")
	}
	if r, _ := ToUint64(f64); r != uint64(f64) {
		t.Fatal("float64 ToUint64 error")
	}
	if r, _ := ToUint64(f64); r != uint64(f64) {
		t.Fatal("*float64 ToUint64 error")
	}
	if r, _ := ToUint64(s1); strconv.FormatUint(r, 10) != s1 {
		t.Fatal("string ToUint64 error")
	}
	if r, _ := ToUint64(&s1); strconv.FormatUint(r, 10) != s1 {
		t.Fatal("*string ToUint64 error")
	}
	if _, err1 := ToUint64(s2); err1 != nil {
		t.Log("err1", err1)
	} else {
		t.Fatal("invalid string ToUint64 error")
	}
	if _, err2 := ToUint64(&s2); err2 != nil {
		t.Log("err2", err2)
	} else {
		t.Fatal("invalid *string ToUint64 error")
	}
	if _, err3 := ToUint64(struct1); err3 != nil {
		t.Log("err3", err3)
	} else {
		t.Fatal("struct ToUint64 error")
	}
	if _, err4 := ToUint64(&struct1); err4 != nil {
		t.Log("err4", err4)
	} else {
		t.Fatal("*struct ToUint64 error")
	}
	if r, err5 := ToUint64(nil1); r != 0 || err5 != nil {
		t.Fatal("nil *struct ToUint64 error")
	}
	if r, err6 := ToUint64(nil); r != 0 || err6 != nil {
		t.Fatal("nil interface{} ToUint64 error")
	}
}

func TestToUint32(t *testing.T) {
	if _, err1 := ToUint32("hello"); err1 != nil {
		t.Log("err1", err1)
	} else {
		t.Fatal("invalid string ToUint32 error")
	}
	if r, _ := ToUint32("3"); r != 3 {
		t.Fatal("string 3 ToUint32 error")
	}
}

func TestToInt64(t *testing.T) {
	if _, err1 := ToInt64("hello"); err1 != nil {
		t.Log("err1", err1)
	} else {
		t.Fatal("invalid string ToInt64 error")
	}
	if r, _ := ToInt64("3"); r != 3 {
		t.Fatal("string 3 ToInt64 error")
	}
}

func TestToInt32(t *testing.T) {
	if _, err1 := ToInt32("hello"); err1 != nil {
		t.Log("err1", err1)
	} else {
		t.Fatal("invalid string ToInt32 error")
	}
	if r, _ := ToInt32("3"); r != 3 {
		t.Fatal("string 3 ToInt32 error")
	}
}

func TestToFloat64(t *testing.T) {
	i, i8, i16, i32, i64 := int(1), int8(2), int16(3), int32(4), int64(5)
	ui, ui8, ui16, ui32, ui64 := uint(1), uint8(2), uint16(3), uint32(4), uint64(5)
	f32, f64 := float32(1.1), float64(2.2)
	s1, s2 := "1.1", "hello"
	struct1 := struct{}{}
	var nil1 *struct{}
	if r, _ := ToFloat64(i); r != float64(i) {
		t.Fatal("int ToFloat64 error")
	}
	if r, _ := ToFloat64(&i); r != float64(i) {
		t.Fatal("*int ToFloat64 error")
	}
	if r, _ := ToFloat64(i8); r != float64(i8) {
		t.Fatal("int8 ToFloat64 error")
	}
	if r, _ := ToFloat64(&i8); r != float64(i8) {
		t.Fatal("*int8 ToFloat64 error")
	}
	if r, _ := ToFloat64(i16); r != float64(i16) {
		t.Fatal("int16 ToFloat64 error")
	}
	if r, _ := ToFloat64(&i16); r != float64(i16) {
		t.Fatal("*int16 ToFloat64 error")
	}
	if r, _ := ToFloat64(i32); r != float64(i32) {
		t.Fatal("int32 ToFloat64 error")
	}
	if r, _ := ToFloat64(&i32); r != float64(i32) {
		t.Fatal("*int32 ToFloat64 error")
	}
	if r, _ := ToFloat64(i64); r != float64(i64) {
		t.Fatal("int64 ToFloat64 error")
	}
	if r, _ := ToFloat64(&i64); r != float64(i64) {
		t.Fatal("*int64 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui); r != float64(ui) {
		t.Fatal("uint ToFloat64 error")
	}
	if r, _ := ToFloat64(ui); r != float64(ui) {
		t.Fatal("*uint ToFloat64 error")
	}
	if r, _ := ToFloat64(ui8); r != float64(ui8) {
		t.Fatal("uint8 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui8); r != float64(ui8) {
		t.Fatal("*uint8 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui16); r != float64(ui16) {
		t.Fatal("uint16 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui16); r != float64(ui16) {
		t.Fatal("*uint16 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui32); r != float64(ui32) {
		t.Fatal("uint32 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui32); r != float64(ui32) {
		t.Fatal("*uint32 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui64); r != float64(ui64) {
		t.Fatal("uint64 ToFloat64 error")
	}
	if r, _ := ToFloat64(ui64); r != float64(ui64) {
		t.Fatal("*uint64 ToFloat64 error")
	}
	if r, _ := ToFloat64(f32); r != float64(f32) {
		t.Fatal("float32 ToFloat64 error")
	}
	if r, _ := ToFloat64(f32); r != float64(f32) {
		t.Fatal("*float32 ToFloat64 error")
	}
	if r, _ := ToFloat64(f64); r != float64(f64) {
		t.Fatal("float64 ToFloat64 error")
	}
	if r, _ := ToFloat64(f64); r != float64(f64) {
		t.Fatal("*float64 ToFloat64 error")
	}
	if r, _ := ToFloat64(s1); strconv.FormatFloat(r, 'f', -1, 64) != s1 {
		t.Fatal("string ToFloat64 error")
	}
	if r, _ := ToFloat64(&s1); strconv.FormatFloat(r, 'f', -1, 64) != s1 {
		t.Fatal("*string ToFloat64 error")
	}
	if _, err1 := ToFloat64(s2); err1 != nil {
		t.Log("err1", err1)
	} else {
		t.Fatal("invalid string ToFloat64 error")
	}
	if _, err2 := ToFloat64(&s2); err2 != nil {
		t.Log("err2", err2)
	} else {
		t.Fatal("invalid *string ToFloat64 error")
	}
	if _, err3 := ToFloat64(struct1); err3 != nil {
		t.Log("err3", err3)
	} else {
		t.Fatal("struct ToFloat64 error")
	}
	if _, err4 := ToFloat64(&struct1); err4 != nil {
		t.Log("err4", err4)
	} else {
		t.Fatal("*struct ToFloat64 error")
	}
	if r, err5 := ToFloat64(nil1); r != 0 || err5 != nil {
		t.Fatal("nil *struct ToFloat64 error")
	}
	if r, err6 := ToFloat64(nil); r != 0 || err6 != nil {
		t.Fatal("nil interface{} ToFloat64 error")
	}
}

func TestToFloat32(t *testing.T) {
	if _, err1 := ToFloat32("hello"); err1 != nil {
		t.Log("err1", err1)
	} else {
		t.Fatal("invalid string ToFloat32 error")
	}
	if r, _ := ToFloat32("3.3"); r != 3.3 {
		t.Fatal("string 3.3 ToFloat32 error")
	}
}
