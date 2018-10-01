package util

import (
	"testing"
)

func TestNewUUIDStr(t *testing.T) {
	uuid := NewUUIDStr()
	if len(uuid) != 36 {
		t.Fatal("uuid check len error")
	}
	t.Log("uuid", uuid)
}

func TestIsInvalidUUID(t *testing.T) {
	if !IsInvalidUUID("") {
		t.Fatal("check empty uuid error")
	}
	if !IsInvalidUUID("123") {
		t.Fatal("check invalid uuid 123 error")
	}
	if IsInvalidUUID("339dc43e-d105-472c-bb6f-e47b6d872051") {
		t.Fatal("check uuid error")
	}
}
