package util

import (
	"os"
	"strconv"
	"time"
)

func GetEnvStr(key string, def string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return def
}

func GetEnvInt(key string, def int) int {
	str := GetEnvStr(key, "")
	if str == "" {
		return def
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return int(i)
}

func GetEnvDuration(key string, def time.Duration) time.Duration {
	str := GetEnvStr(key, "")
	if str == "" {
		return def
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		return def
	}
	return duration
}
