package main

import (
	"log"
	"os"

	"github.com/MaxRazen/tutor/internal/config"
	"github.com/MaxRazen/tutor/internal/routes"
	"github.com/MaxRazen/tutor/pkg/google"
	"github.com/MaxRazen/tutor/pkg/oauth"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
)

type ServerConfig struct {
	Port      string
	Host      string
	WatchMode bool
}

func InitServer(cfg config.RuntimeConfig) {
	server := fiber.New()
	routes.SetRootTemplate(&publicRoot)
	authMiddleware := routes.NewAuthMiddleware()

	compressConfig := compress.Config{
		Level: compress.LevelDefault,
	}

	server.Use(compress.New(compressConfig))

	filesystemConfig := filesystem.Config{
		Root:       content,
		PathPrefix: "ui/public",
		Browse:     true,
		MaxAge:     60 * 60 * 24 * 7, // 7days
	}

	if cfg.DevMode {
		filesystemConfig = filesystem.Config{
			Root:   os.DirFS("ui/public"),
			Browse: true,
		}
	}

	server.Use("/assets", filesystem.New(filesystemConfig))

	server.Get("/auth/redirect/:provider", routes.AuthRedirect())
	server.Get("/auth/callback/:provider", routes.AuthCallback())

	server.Get("*", routes.HomeHandler(), authMiddleware)

	err := server.Listen(cfg.GetServerHost())

	if err != nil {
		log.Fatalln(err.Error())
	}
}

func InitOAuthProviders() {
	provider := google.New(
		config.GetEnv(config.GOOGLE_OAUTH_CLIENT_ID, ""),
		config.GetEnv(config.GOOGLE_OAUTH_SECRET, ""),
		config.GetEnv(config.GOOGLE_OAUTH_CALLBACK_URL, ""),
		[]string{},
	)
	oauth.UseProvider(provider)
}
