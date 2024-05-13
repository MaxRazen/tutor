package main

import (
	"embed"
	"os"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/config"
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
	cloud.PrepareClient(credentials, "credentials/gcp.json")

	InitOAuthProviders()

	InitServer(config.NewConfig(mode, os.Args[1:]))
}
