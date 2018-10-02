package grpcutil

import (
	"testing"

	"google.golang.org/grpc"
)

func TestNewServer(t *testing.T) {
	t.Log(NewServer(func(server *grpc.Server) {
		t.Log("register service")
	}))
}
