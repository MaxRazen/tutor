package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/ui"
	"github.com/MaxRazen/tutor/pkg/oauth"
	fiber "github.com/gofiber/fiber/v3"
)

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

func HomeHandler() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
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

func AuthRedirect() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		providerName := c.Params("provider", "")
		if providerName == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		provider, err := oauth.GetProvider(providerName)

		if err != nil {
			log.Println(err.Error())
			return c.SendStatus(http.StatusBadRequest)
		}

		authUrl := provider.BeginAuth("")

		return c.Redirect().To(authUrl)
	}
}

func AuthCallback() func(c fiber.Ctx) error {
	return func(c fiber.Ctx) error {
		providerName := c.Params("provider", "")

		if providerName == "" {
			return c.SendStatus(http.StatusNotFound)
		}

		provider, err := oauth.GetProvider(providerName)

		if err != nil {
			log.Println(err.Error())
			return c.SendStatus(http.StatusBadRequest)
		}

		user, err := provider.CompleteAuth(c.Queries())

		if err != nil {
			log.Println(err.Error())

			data, _ := ui.WrapWithKey(ui.NewAlert(ui.AlertError, ui.AlertMessageAuthenticationFailed), "alert")

			return respondWithHtml(c, data)
		}

		// TODO: Save user to DB

		accessToken, err := auth.SignAccessToken(user)

		if err != nil {
			log.Println(err.Error())

			data, _ := ui.WrapWithKey(ui.NewAlert(ui.AlertError, ui.AlertMessageAuthenticationFailed), "alert")

			return respondWithHtml(c, data)
		}

		c.Cookie(auth.CreateAccessTokenCookie(accessToken))

		data, err := ui.WrapUserInfo(user, accessToken)

		if err != nil {
			log.Println(err.Error())
		}

		return respondWithHtml(c, data)
	}
}

func respondWithHtml(c fiber.Ctx, data any) error {
	c.Context().Response.Header.SetStatusCode(http.StatusOK)
	c.Context().Response.Header.Add("Content-Type", "text/html")

	return rootTemplate.Execute(c.Context().Response.BodyWriter(), data)
}
