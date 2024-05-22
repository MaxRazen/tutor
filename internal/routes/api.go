package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/room"
	fiber "github.com/gofiber/fiber/v2"
)

func CreateRoomHandler() routeHandler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(int)

		user, err := auth.FindUser(userId)

		if err != nil {
			log.Printf("api/room: %v", err)
			return respondWithUnauthorizedError(c)
		}

		var data room.CreationData
		if err := json.Unmarshal(c.Body(), &data); err != nil {
			log.Printf("api/room: %v", err)
			return c.SendStatus(http.StatusBadRequest)
		}

		r, err := room.CreateRoom(data, user)

		if err != nil {
			log.Printf("api/room: %v", err)
			return c.SendStatus(http.StatusInternalServerError)
		}

		serializableData := fiber.Map{
			"roomId":    r.ID,
			"expiresIn": time.Now().Add(time.Hour * 12).Unix(),
		}

		return c.JSON(serializableData)
	}
}
