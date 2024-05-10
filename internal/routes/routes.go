package routes

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/MaxRazen/tutor/internal/ui"
	"github.com/MaxRazen/tutor/pkg/oauth"
	fiber "github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
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
			c.Context().NotFound()
			return nil
		}

		m := make(map[string]any)

		data, _ := ui.NewTemplateData(m)

		return responseWithHtml(c, data)
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
		// TODO: Handle all nil returns
		providerName := c.Params("provider", "")

		if providerName == "" {
			return c.SendStatus(http.StatusBadRequest)
		}

		provider, err := oauth.GetProvider(providerName)

		if err != nil {
			log.Println(err.Error())
			return c.SendStatus(http.StatusBadRequest)
		}

		user, err := provider.CompleteAuth(c.Queries())

		if err != nil {
			log.Println(err.Error())
			return nil
		}

		// TODO: Save user to DB

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"uid": user.SocialID,
			"exp": time.Now().Add(time.Minute * 1),
		})
		accessToken, err := token.SignedString([]byte("some_secret_key"))

		if err != nil {
			log.Println(err.Error())
			return nil
		}

		data, err := ui.WrapUserInfo(user, accessToken)

		if err != nil {
			log.Println(err.Error())
		}

		return responseWithHtml(c, data)
	}
}

func responseWithHtml(c fiber.Ctx, data any) error {
	c.Context().Response.Header.SetStatusCode(http.StatusOK)
	c.Context().Response.Header.Add("Content-Type", "text/html")

	return rootTemplate.Execute(c.Context().Response.BodyWriter(), data)
}
