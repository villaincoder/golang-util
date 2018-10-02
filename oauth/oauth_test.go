package oauth

import (
	"testing"
	"time"

	"gopkg.in/go-oauth2/redis.v1"
)

func TestLoadEnvServerConfig(t *testing.T) {
	config := LoadEnvConfig(&Config{
		RedisConfig: &redis.Config{
			DB: 11,
		},
		AccessTokenExp: time.Hour * 1,
	})
	if config.RedisConfig.DB != 11 {
		t.Fatal("redis config db error")
	}
	if config.AccessTokenExp != time.Hour*1 {
		t.Fatal("config access token exp error")
	}
	if LoadEnvConfig(nil) == nil {
		t.Fatal("load env config with nil error")
	}
}
