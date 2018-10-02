package util

import "testing"

func TestStr2Ptr(t *testing.T) {
	ptr1 := Str2Ptr("")
	if ptr1 != nil {
		t.Fatal("empty string to pointer error")
	}
	ptr2 := Str2Ptr("1")
	if ptr2 == nil {
		t.Fatal("non-empty string to pointer error")
	}
}

func TestPtr2Str(t *testing.T) {
	str1 := Ptr2Str(nil)
	if str1 != "" {
		t.Fatal("nil to string error")
	}
	s := "1"
	str2 := Ptr2Str(&s)
	if str2 != s {
		t.Fatal("pointer to string error")
	}
}

func TestStrs2Ptrs(t *testing.T) {
	ptrs1 := Strs2Ptrs(nil)
	if ptrs1 != nil {
		t.Fatal("nil strings to pointers error")
	}
	ptrs2 := Strs2Ptrs([]string{"", "1"})
	if ptrs2 == nil {
		t.Fatal("strings to pointers error")
	}
	if ptrs2[0] != nil {
		t.Fatal("empty string item to pointer item error")
	}
	if ptrs2[1] == nil {
		t.Fatal("non-empty string item to pointer item error")
	}
}

func TestPtrs2Strs(t *testing.T) {
	strs1 := Ptrs2Strs(nil)
	if strs1 != nil {
		t.Fatal("nil pointers to strings error")
	}
	s := "1"
	strs2 := Ptrs2Strs([]*string{nil, &s})
	if strs2 == nil {
		t.Fatal("pointers to strings error")
	}
	if strs2[0] != "" {
		t.Fatal("nil pointer item to string item error")
	}
	if strs2[1] != s {
		t.Fatal("non-nil pointer item to string item error")
	}
}

func TestStringFallback(t *testing.T) {
	if StrFallback("", "123") != "123" {
		t.Fatal("get empty string fallback error")
	}
	if StrFallback("hello", "123") != "hello" {
		t.Fatal("get non-empty string fallback error")
	}
}
