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

var mode string

func main() {
	InitServer(config.NewConfig(mode, os.Args[1:]))
}
