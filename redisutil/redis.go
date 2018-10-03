package redisutil

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

func LoadEnvConfig(config *redis.Options) *redis.Options {
	if config == nil {
		config = &redis.Options{}
	}
	config.Addr = util.GetEnvStr("REDIS_ADDR", util.StrFallback(config.Addr, "127.0.0.1:6379"))
	config.Password = util.GetEnvStr("REDIS_PASSWORD", util.StrFallback(config.Password, ""))
	config.DB = util.GetEnvInt("REDIS_DB", config.DB)
	return config
}

func OpenRedis(config *redis.Options) (client *redis.Client, err error) {
	client = redis.NewClient(config)
	err = errors.WithStack(client.Ping().Err())
	if err != nil {
		client.Close()
		client = nil
	}
	return
}
