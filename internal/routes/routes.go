package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/MaxRazen/tutor/internal/room"
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

func ShowRoomHandler() routeHandler {
	return func(c *fiber.Ctx) error {
		roomId, err := strconv.Atoi(c.Params("id"))
		userId := c.Locals("userId").(int)

		if err != nil || roomId == 0 {
			return c.SendStatus(http.StatusNotFound)
		}

		roomRecord, err := room.FindRoom(roomId, userId)

		if err != nil || roomRecord == nil {
			log.Println(err)
			return c.SendStatus(http.StatusNotFound)
		}

		m := make(map[string]any)
		m["roomId"] = roomId

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
