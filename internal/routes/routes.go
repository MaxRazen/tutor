package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	fiber "github.com/gofiber/fiber/v3"
)

var spaPaths = [...]string{
	"/",
	"/login",
	"/about",
}

func HomeHandler(publicRoot *embed.FS) func(c fiber.Ctx) error {
	layout, err := template.ParseFS(publicRoot, "ui/templates/index.html")

	if err != nil {
		log.Fatalf("template loading failed: %v", err)
	}

	return func(c fiber.Ctx) error {
		path := string(c.Context().Request.URI().Path())

		if !isPathSupported(path) {
			c.Context().NotFound()
			return nil
		}

		c.Context().Response.Header.SetStatusCode(http.StatusOK)
		c.Context().Response.Header.Add("Content-Type", "text/html")

		return layout.Execute(c.Context().Response.BodyWriter(), nil)
	}
}

func isPathSupported(requestedPath string) bool {
	for _, path := range spaPaths {
		if requestedPath == path {
			return true
		}
	}
	return false
}

func AuthRedirect() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return c.Redirect().To("/")
	}
}

func AuthCallback() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		return nil
	}
}
