// Package config handles all the necessary configuration from env variables
package config

import (
	"log"
	"os"
)

type Config struct {
	SshAdminPassword string
}

func Load() Config {
	return Config{
		SshAdminPassword: mustGetEnv("SSH_ADMIN_PASSWORD"),
	}
}

func mustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		log.Fatalf("missing required environment variable: %s", key)
	}

	return val
}

