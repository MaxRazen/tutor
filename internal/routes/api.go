package routes

import (
	fiber "github.com/gofiber/fiber/v2"
)

func CreateRoomHandler() routeHandler {
	return func(c *fiber.Ctx) error {
		room := make(map[string]string)
		room["id"] = "1001"

		return c.JSON(room)
	}
}
