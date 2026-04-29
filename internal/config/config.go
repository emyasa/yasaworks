// Package config handles all the necessary configuration from env variables
package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	SshAdminPassword string
	AllowAnonymous bool
}

func Load() Config {
	return Config{
		SshAdminPassword: mustGetEnv("SSH_ADMIN_PASSWORD"),
		AllowAnonymous: getEnvBool("ALLOW_ANONYMOUS", false),
	}
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}

	return val
}

func getEnvBool(key string, defaultVal bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		return defaultVal
	}

	parsed, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("invalid boolean value for %s: %s", key, val)
	}

	return parsed
}

