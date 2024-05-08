package main

import (
	"embed"
	"html/template"
	"net/http"
	"os"

	fiber "github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/filesystem"
)

//go:embed templates/index.html
var publicRoot embed.FS

//go:embed public
var content embed.FS

var mode string

const (
	MODE_WATCH   = "watch"
	MODE_RELEASE = "release"
)

type ServerConfig struct {
	Port      string
	Host      string
	WatchMode bool
}

var supportedUiPaths = [...]string{
	"/",
	"/login",
	"/about",
}

func main() {
	if mode == "" {
		mode = MODE_RELEASE
	}

	serverConfig := ServerConfig{
		Port:      "3000",
		Host:      "",
		WatchMode: mode == MODE_WATCH,
	}

	layout, err := template.ParseFS(publicRoot, "templates/index.html")

	if err != nil {
		println(err.Error())
		os.Exit(1)
		return
	}

	server := fiber.New()

	compressConfig := compress.Config{
		Level: compress.LevelDefault,
	}

	if serverConfig.WatchMode {
		compressConfig.Level = compress.LevelBestSpeed
	}
	server.Use(compress.New(compressConfig))

	filesystemConfig := filesystem.Config{
		Root:       content,
		PathPrefix: "public",
		Browse:     true,
		MaxAge:     60 * 60 * 24 * 7, // 7days
	}

	if serverConfig.WatchMode {
		filesystemConfig = filesystem.Config{
			Root:   os.DirFS("public"),
			Browse: true,
		}
	}

	server.Use("/assets", filesystem.New(filesystemConfig))

	server.Get("/api/v1/ping", func(c fiber.Ctx) error {
		return c.SendString("pong")
	})

	server.Get("*", func(c fiber.Ctx) error {
		path := string(c.Context().Request.URI().Path())

		if !isPathSupported(path) {
			c.Context().NotFound()
			return nil
		}

		c.Context().Response.Header.SetStatusCode(http.StatusOK)
		c.Context().Response.Header.Add("Content-Type", "text/html")

		return layout.Execute(c.Context().Response.BodyWriter(), serverConfig)
	})

	server.Listen(serverConfig.Host + ":" + serverConfig.Port)
}

func isPathSupported(requestedPath string) bool {
	for _, path := range supportedUiPaths {
		if requestedPath == path {
			return true
		}
	}
	return false
}
