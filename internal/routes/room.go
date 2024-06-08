package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/MaxRazen/tutor/internal/auth"
	"github.com/MaxRazen/tutor/internal/room"
	"github.com/MaxRazen/tutor/internal/ui"
	fiber "github.com/gofiber/fiber/v2"
)

func RoomCreateHandler() routeHandler {
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

func RoomShowHandler() routeHandler {
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

		roomHistory, err := room.LoadRoomHistory(roomId)

		if err != nil {
			log.Println("routes/room:", err)
			return c.SendStatus(http.StatusInternalServerError)
		}

		m := make(map[string]any)
		m["roomId"] = roomId
		m["history"] = roomHistory

		data, _ := ui.NewTemplateData(m)

		return respondWithHtml(c, data)
	}
}
