package grpcutil

import (
	"testing"

	"google.golang.org/grpc"
)

func TestLoadEnvConfig(t *testing.T) {
	config := LoadEnvConfig(&Config{
		Port: "1234",
	})
	if config.Port != "1234" {
		t.Fatal("config port error")
	}
	if LoadEnvConfig(nil) == nil {
		t.Fatal("load env config with nil error")
	}
}

func TestNewServer(t *testing.T) {
	NewServer(func(server *grpc.Server) {})
}
