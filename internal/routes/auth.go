package routes

import (
	"log"
	"net/http"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/ui"
	"github.com/MaxRazen/tutor/pkg/oauth"
	fiber "github.com/gofiber/fiber/v2"
)

func AuthRedirect() routeHandler {
	return func(c *fiber.Ctx) error {
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

		return c.Redirect(authUrl)
	}
}

func AuthCallback() routeHandler {
	return func(c *fiber.Ctx) error {
		providerName := c.Params("provider", "")

		if providerName == "" {
			return c.SendStatus(http.StatusNotFound)
		}

		provider, err := oauth.GetProvider(providerName)

		if err != nil {
			log.Println(err.Error())
			return c.SendStatus(http.StatusBadRequest)
		}

		// TODO: the first argument `token` can be used to refresh accessToken
		_, profile, err := provider.CompleteAuth(c.Queries())

		if err != nil {
			return respondWithAuthError(c, err)
		}

		user, err := auth.FindOrCreateUser(profile)

		if err != nil {
			return respondWithAuthError(c, err)
		}

		accessToken, err := auth.IssueAccessToken(user)

		if err != nil {
			return respondWithAuthError(c, err)
		}

		c.Cookie(auth.CreateAccessTokenCookie(accessToken))

		data, err := ui.WrapUserInfo(user, accessToken)

		if err != nil {
			log.Println(err.Error())
		}

		return respondWithHtml(c, data)
	}
}

func AuthLogout() routeHandler {
	return func(c *fiber.Ctx) error {
		c.Cookie(auth.ExpireAccessTokenCookie())

		return c.Redirect("/login")
	}
}

func respondWithAuthError(c *fiber.Ctx, err error, alertMessage ...string) error {
	log.Println(err.Error())

	msg := ui.AlertMessageAuthenticationFailed
	if len(alertMessage) > 0 {
		msg = alertMessage[0]
	}

	data, _ := ui.WrapWithKey(ui.NewAlert(ui.AlertError, msg), "alert")

	return respondWithHtml(c, data)
}
