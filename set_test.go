package util

import "testing"

func TestNewSet(t *testing.T) {
	set1 := NewSet()
	if set1.Size() != 0 {
		t.Fatal("empty item new set error")
	}
	set2 := NewSet("1", "2", "3")
	if set2.Size() != 3 {
		t.Fatal("item new set error")
	}
	set3 := NewSet("1", "2", "3", "1")
	if set3.Size() != 3 {
		t.Fatal("duplicate item new set error")
	}
}

func TestSet(t *testing.T) {
	set := NewSet()
	set.Add("1")
	if set.Size() != 1 {
		t.Fatal("add 1 check size error")
	}
	if !set.Contains("1") {
		t.Fatal("add 1 check contains error")
	}
	set.Add("2", "3")
	if set.Size() != 3 {
		t.Fatal("add 2,3 check size error")
	}
	if !set.Contains("2") {
		t.Fatal("add 2 check contains error")
	}
	if !set.Contains("3") {
		t.Fatal("add 3 check contains error")
	}
	set.Remove("1")
	if set.Size() != 2 {
		t.Fatal("remove 1 check size error")
	}
	if set.Contains("1") {
		t.Fatal("remove 1 check contains error")
	}
	set.Clear()
	if set.Size() != 0 {
		t.Fatal("clear check size error")
	}
}

func TestSet_Foreach(t *testing.T) {
	set := NewSet("1", "2", "3")
	s1 := ""
	set.Foreach(func(i interface{}) bool {
		s1 += i.(string)
		return false
	})
	if len(s1) != 3 {
		t.Fatal("foreach no break error")
	}
	s2 := ""
	set.Foreach(func(i interface{}) bool {
		s2 += i.(string)
		return true
	})
	if len(s2) != 1 {
		t.Fatal("foreach break error")
	}
}
