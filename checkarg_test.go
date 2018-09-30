package util

import (
	"testing"
	"time"
)

func TestCheckArgs(t *testing.T) {
	err1 := CheckArgs(map[string]interface{}{
		"key": "",
	})
	if err1 == nil {
		t.Fatal("check empty string error")
	}
	t.Log("err1", err1)
	err1 = CheckArgs(map[string]interface{}{
		"key": "value",
	})
	if err1 != nil {
		t.Fatal("check non-empty string error")
	}

	err2 := CheckArgs(map[string]interface{}{
		"key": time.Time{},
	})
	if err2 == nil {
		t.Fatal("check invalid time error")
	}
	t.Log("err2", err2)
	err2 = CheckArgs(map[string]interface{}{
		"key": time.Now(),
	})
	if err2 != nil {
		t.Fatal("check non-invalid time error")
	}

	err3 := CheckArgs(map[string]interface{}{
		"key": nil,
	})
	if err3 == nil {
		t.Fatal("check nil value error")
	}
	t.Log("err3", err3)
	err3 = CheckArgs(map[string]interface{}{
		"key": &struct{}{},
	})
	if err3 != nil {
		t.Fatal("check non-nil value error")
	}
}

func TestCheckStrArrayArg(t *testing.T) {
	err1 := CheckStrArrayArg("field", []string{
		"", "t1", "t2",
	})
	if err1 == nil {
		t.Fatal("check empty item error")
	}
	t.Log("err1", err1)
	err1 = CheckStrArrayArg("field", []string{
		"t1", "t2", "t3",
	})
	if err1 != nil {
		t.Fatal("check non-empty item error")
	}

	err2 := CheckStrArrayArg("field", nil)
	if err2 == nil {
		t.Fatal("check nil array error")
	}
	t.Log("err2", err2)
}
