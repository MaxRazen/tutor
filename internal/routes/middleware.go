package routes

import (
	"log"
	"net/http"
	"strings"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/utils"
	fiber "github.com/gofiber/fiber/v2"
)

func AuthMiddleware() routeHandler {
	filteredPaths := []string{
		"/login",
	}

	return func(c *fiber.Ctx) error {
		path := c.Path()

		if utils.InSlice(path, filteredPaths) {
			return c.Next()
		}

		accessToken := c.Cookies("jwt", "")

		if accessToken == "" {
			return respondWithUnauthorizedError(c)
		}

		ok, claims := auth.ValidateToken(accessToken)

		if !ok {
			return respondWithUnauthorizedError(c)
		}

		claimUID, ok := claims["uid"]

		if !ok {
			log.Printf("uid is not present in jwt token")
			return respondWithUnauthorizedError(c)
		}

		userId, ok := claimUID.(float64)

		if !ok {
			log.Printf("userId has invalid value %v(%T)", claimUID, claimUID)
			return respondWithUnauthorizedError(c)
		}

		c.Locals("userId", int(userId))

		return c.Next()
	}
}

func respondWithUnauthorizedError(c *fiber.Ctx) error {
	if strings.HasPrefix(c.Path(), "/api") {
		return unauthorized(c)
	}
	return c.Redirect("/login")
}

func unauthorized(c *fiber.Ctx) error {
	c.Response().SetStatusCode(http.StatusUnauthorized)

	return c.JSON(map[string]string{
		"message": "unauthorized",
	})
}
