package grpc

import (
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

type ServeConfig struct {
	Port string
}

func LoadEnvServeConfig(config *ServeConfig) *ServeConfig {
	if config == nil {
		config = &ServeConfig{}
	}
	config.Port = util.GetEnv("DB_PORT", util.StringFallback(config.Port, "50051"))
	return config
}

type ServiceRegisterFunc func(server *grpc.Server)

func NewServer(serviceRegister ServiceRegisterFunc) (server *grpc.Server) {
	server = grpc.NewServer()
	serviceRegister(server)
	reflection.Register(server)
	return
}

func Serve(config *ServeConfig, server *grpc.Server) (err error) {
	listener, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		return
	}
	if err = server.Serve(listener); err != nil {
		return
	}
	return
}
