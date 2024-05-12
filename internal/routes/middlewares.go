package routes

import (
	"net/http"
	"strings"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/utils"
	fiber "github.com/gofiber/fiber/v2"
)

func AuthMiddleware() routeHandler {
	protectedPaths := []string{
		"/",
		"/about",
		"/api/v1/room",
	}

	return func(c *fiber.Ctx) error {
		path := c.Path()

		if !utils.InSlice(path, protectedPaths) {
			return c.Next()
		}

		accessToken := c.Cookies("jwt", "")

		if accessToken == "" || !auth.ValidateToken(accessToken) {
			if strings.HasPrefix(path, "/api") {
				return unauthorized(c)
			}
			return c.Redirect("/login")
		}

		return c.Next()
	}
}

func unauthorized(c *fiber.Ctx) error {
	c.Response().SetStatusCode(http.StatusUnauthorized)

	return c.JSON(map[string]string{
		"message": "unauthorized",
	})
}
