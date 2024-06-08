package main

import (
	"embed"
	"os"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/config"
	"github.com/MaxRazen/tutor/internal/db"
	"github.com/MaxRazen/tutor/pkg/memcache"
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
	runtimeConfig := config.NewConfig(mode, os.Args[1:])
	auth.SetSecretKey()
	db.Connect()
	cloud.PrepareClient(credentials, "credentials/gcp.json")
	memcache.Init(memcache.NewGCPStorage(cloud.GetBucket(), "memcache.db"))

	if runtimeConfig.AutoMigrate {
		db.MigrateDB()
	}

	InitOAuthProviders()

	InitServer(runtimeConfig)
}
