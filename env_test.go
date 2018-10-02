package util

import (
	"testing"
	"time"
)

func TestGetEnv(t *testing.T) {
	if GetEnvStr("GOROOT", "") == "" {
		t.Fatal("get exists env error")
	}
	if GetEnvStr(NewUUIDStr(), "1234") != "1234" {
		t.Fatal("get not exists env error")
	}
}

func TestGetEnvInt(t *testing.T) {
	if GetEnvInt(NewUUIDStr(), 10) != 10 {
		t.Fatal("get not exists int env error")
	}
}

func TestGetEnvDuration(t *testing.T) {
	if GetEnvDuration(NewUUIDStr(), time.Hour*10) != time.Hour*10 {
		t.Fatal("get not exists duration env error")
	}
}
