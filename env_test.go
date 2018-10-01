package util

import (
	"testing"
)

func TestGetEnv(t *testing.T) {
	if GetEnv("GOROOT", "") == "" {
		t.Fatal("get exists env error")
	}
	if GetEnv(NewUUIDStr(), "1234") != "1234" {
		t.Fatal("get not exists env error")
	}
}
