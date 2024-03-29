package config

import (
	"os"
	"strconv"
)

type Flag struct {
	ConfigFile string
}

var ConfigFlag Flag

func GetEnv(key string, defVal string) string {
	if env, ok := os.LookupEnv(key); ok {
		return env
	}

	return defVal

}

func GetEnvInt(key string, defVal int64) int64 {
	if env, ok := os.LookupEnv(key); ok {
		envInt, err := strconv.ParseInt(env, 10, 64)
		if err == nil {
			return envInt
		}
	}

	return defVal

}
