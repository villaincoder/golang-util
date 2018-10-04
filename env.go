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

func GetEnvInt64(key string, def int64) int64 {
	str := GetEnvStr(key, "")
	if str == "" {
		return def
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return def
	}
	return i
}

func GetEnvUint64(key string, def uint64) uint64 {
	str := GetEnvStr(key, "")
	if str == "" {
		return def
	}
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return def
	}
	return i
}

func GetEnvInt(key string, def int) int {
	return int(GetEnvInt64(key, int64(def)))
}
func GetEnvUint(key string, def uint) uint {
	return uint(GetEnvUint64(key, uint64(def)))
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
