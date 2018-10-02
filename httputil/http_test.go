package httputil

import (
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestNewRouter(t *testing.T) {
	router := NewRouter(func(router *httprouter.Router) {
		t.Log("register route")
	})
	t.Log(router)
}
