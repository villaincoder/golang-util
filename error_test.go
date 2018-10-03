package util

import (
	"errors"
	"testing"
)

func TestPrintError(t *testing.T) {
	PrintError(errors.New("test error"))
}
