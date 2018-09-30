package util

import "github.com/satori/go.uuid"

func NewUUID() *string {
	u := uuid.NewV4().String()
	return &u
}
