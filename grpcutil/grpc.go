package grpcutil

import (
	"github.com/pkg/errors"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type Config struct {
	Port string
}

func LoadEnvConfig(config *Config) *Config {
	if config == nil {
		config = &Config{}
	}
	config.Port = util.GetEnvStr("GRPC_PORT", util.StrFallback(config.Port, "50051"))
	return config
}

type RegisterServiceFunc func(server *grpc.Server)

func NewServer(registerService RegisterServiceFunc) (server *grpc.Server) {
	server = grpc.NewServer()
	registerService(server)
	reflection.Register(server)
	return
}

func Serve(config *Config, server *grpc.Server) (err error) {
	listener, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	if err = server.Serve(listener); err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
