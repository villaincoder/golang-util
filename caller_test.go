package util

import (
	"math"
	"strings"
	"testing"
)

func TestGetCallerName(t *testing.T) {
	name1 := GetCallerName(math.MaxInt32)
	if name1 != "???" {
		t.Fatal("get invalid caller error")
	}
	name2 := GetCallerName(0)
	if !strings.HasSuffix(name2, "TestGetCallerName") {
		t.Fatal("get caller error")
	}
}
