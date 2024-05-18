package main

import (
	"embed"
	"os"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/config"
	"github.com/MaxRazen/tutor/internal/db"
)

//go:embed ui/templates/index.html
var publicRoot embed.FS

//go:embed ui/public
var content embed.FS

//go:embed credentials
var credentials embed.FS

var mode string

func main() {
	config.LoadEnv(credentials, "credentials/env")
	db.Connect()
	db.MigrateDB()
	cloud.PrepareClient(credentials, "credentials/gcp.json")
	runtimeConfig := config.NewConfig(mode, os.Args[1:])

	InitOAuthProviders()

	InitServer(runtimeConfig)
}
