package util

import "github.com/satori/go.uuid"

func NewUUIDStr() string {
	return uuid.NewV4().String()
}

func IsInvalidUUID(input string) bool {
	if input == "" {
		return true
	}
	if _, err := uuid.FromString(input); err != nil {
		return true
	}
	return false
}
