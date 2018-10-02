package http

import (
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestLoadEnvConfig(t *testing.T) {
	config := LoadEnvConfig(&Config{
		Port: "8888",
	})
	if config.Port != "8888" {
		t.Fatal("config port error")
	}
	if LoadEnvConfig(nil) == nil {
		t.Fatal("load env config with nil error")
	}
}

func TestNewRouter(t *testing.T) {
	router := NewRouter(func(router *httprouter.Router) {
		t.Log("register route")
	})
	t.Log(router)
}
