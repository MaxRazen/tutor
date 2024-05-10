package config

import (
	"embed"
	"log"

	"github.com/joho/godotenv"
)

const (
	GOOGLE_OAUTH_CLIENT_ID    = "GOOGLE_OAUTH_CLIENT_ID"
	GOOGLE_OAUTH_SECRET       = "GOOGLE_OAUTH_SECRET"
	GOOGLE_OAUTH_CALLBACK_URL = "GOOGLE_OAUTH_CALLBACK_URL"
)

var envVariables map[string]string

func LoadEnv(file embed.FS, name string) {
	r, err := file.Open(name)

	if err != nil {
		log.Fatalln("env file cannot be opened:", err.Error())
	}

	m, err := godotenv.Parse(r)

	if err != nil {
		log.Fatalln("env file cannot be parsed:", err.Error())
	}

	envVariables = m
}

func GetEnv(key string, defaultValue string) string {
	v, ok := envVariables[key]
	if ok && v != "" {
		return v
	}
	return defaultValue
}
