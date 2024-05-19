package config

import (
	"embed"
	"log"

	"github.com/joho/godotenv"
)

const (
	APP_KEY                   = "APP_KEY"
	DB_DSN                    = "DB_DSN"
	GOOGLE_OAUTH_CLIENT_ID    = "GOOGLE_OAUTH_CLIENT_ID"
	GOOGLE_OAUTH_SECRET       = "GOOGLE_OAUTH_SECRET"
	GOOGLE_OAUTH_CALLBACK_URL = "GOOGLE_OAUTH_CALLBACK_URL"
	STORAGE_BUCKET_NAME       = "STORAGE_BUCKET_NAME"
)

var envVariables map[string]string

func LoadEnv(credDir embed.FS, name string) {
	b, err := credDir.ReadFile(name)

	if err != nil {
		log.Fatalln("env file cannot be read:", err.Error())
	}

	m, err := godotenv.UnmarshalBytes(b)

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
