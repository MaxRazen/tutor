package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/MaxRazen/tutor/internal/ui"
	fiber "github.com/gofiber/fiber/v2"
)

type routeHandler = fiber.Handler

var spaPaths = [...]string{
	"/",
	"/login",
	"/about",
}

var rootTemplate *template.Template

func SetRootTemplate(publicRoot *embed.FS) {
	layout, err := template.ParseFS(publicRoot, "ui/templates/index.html")

	if err != nil {
		log.Fatalf("template loading failed: %v", err)
	}

	rootTemplate = layout
}

func HomeHandler() routeHandler {
	return func(c *fiber.Ctx) error {
		path := string(c.Context().Request.URI().Path())

		if !isPathSupported(path) {
			return c.SendStatus(http.StatusNotFound)
		}

		m := make(map[string]any)

		data, _ := ui.NewTemplateData(m)

		return respondWithHtml(c, data)
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

func respondWithHtml(c *fiber.Ctx, data any) error {
	c.Context().Response.Header.SetStatusCode(http.StatusOK)
	c.Context().Response.Header.Add("Content-Type", "text/html")

	return rootTemplate.Execute(c.Context().Response.BodyWriter(), data)
}
