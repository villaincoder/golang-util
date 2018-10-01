package grpc

import (
	"testing"

	"google.golang.org/grpc"
)

func TestLoadEnvServeConfig(t *testing.T) {
	config := LoadEnvServeConfig(&ServeConfig{
		Port: "1234",
	})
	if config.Port != "1234" {
		t.Fatal("config port error")
	}
	if LoadEnvServeConfig(nil) == nil {
		t.Fatal("load env config with nil error")
	}
}

func TestNewServer(t *testing.T) {
	NewServer(func(server *grpc.Server) {})
}
