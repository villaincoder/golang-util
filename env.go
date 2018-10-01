package util

import "os"

func GetEnv(key, def string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}
	return def
}
