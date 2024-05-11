package routes

import (
	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/utils"
	"github.com/gofiber/fiber/v3"
)

func NewAuthMiddleware() fiber.Handler {
	protectedPaths := []string{
		"/",
		"/about",
	}

	return func(c fiber.Ctx) error {
		path := c.Path()

		if !utils.InSlice(path, protectedPaths) {
			return c.Next()
		}

		accessToken := c.Cookies("jwt", "")

		if accessToken == "" || !auth.ValidateToken(accessToken) {
			return c.Redirect().To("/login")
		}

		return c.Next()
	}
}
