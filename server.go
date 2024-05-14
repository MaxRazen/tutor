package main

import (
	"log"
	"net/http"

	"github.com/MaxRazen/tutor/internal/config"
	"github.com/MaxRazen/tutor/internal/routes"
	"github.com/MaxRazen/tutor/pkg/google"
	"github.com/MaxRazen/tutor/pkg/oauth"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type ServerConfig struct {
	Port      string
	Host      string
	WatchMode bool
}

func InitServer(cfg config.RuntimeConfig) {
	server := fiber.New()
	routes.SetRootTemplate(&publicRoot)
	authMiddleware := routes.AuthMiddleware()

	compressConfig := compress.Config{
		Level: compress.LevelDefault,
	}

	server.Use(compress.New(compressConfig))

	filesystemConfig := filesystem.Config{
		Root:       http.FS(content),
		PathPrefix: "ui/public",
		Browse:     true,
		MaxAge:     60 * 60 * 24 * 7, // 7days
	}

	if cfg.DevMode {
		filesystemConfig = filesystem.Config{
			Root:   http.Dir("ui/public"),
			Browse: true,
		}
	}

	server.Use("/assets", filesystem.New(filesystemConfig))

	api := server.Group("/api/v1", authMiddleware)
	api.Post("room", routes.CreateRoomHandler())

	server.Use("/ws", authMiddleware, routes.WebsocketMiddleware())
	server.Get("/ws/room/:id", routes.RoomWebsocketHandler())

	server.Get("/auth/redirect/:provider", routes.AuthRedirect())
	server.Get("/auth/callback/:provider", routes.AuthCallback())

	server.Get("*", authMiddleware, routes.HomeHandler())

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
