package main

import (
	"embed"
	"os"

	"github.com/MaxRazen/tutor/internal/config"
)

//go:embed ui/templates/index.html
var publicRoot embed.FS

//go:embed ui/public
var content embed.FS

//go:embed env
var envFile embed.FS

var mode string

func main() {
	config.LoadEnv(envFile, "env")
	InitOAuthProviders()

	InitServer(config.NewConfig(mode, os.Args[1:]))
}
