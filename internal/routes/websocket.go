package routes

import (
	"log"
	"time"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func WebsocketMiddleware() routeHandler {
	return func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			println("ws conn is upgraded by middleware")
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	}
}

func RoomWebsocketHandler() routeHandler {
	return websocket.New(func(c *websocket.Conn) {
		if c.Locals("allowed") == false {
			c.Close()
			return
		}
		defer func() {
			println("connection per room", c.Params("id"), "is closed")
		}()
		// TODO: Fetch Room metadata c.Params("id")

		messages := make(chan []byte)
		go func() {
			defer close(messages)

			for {
				_, msg, err := c.ReadMessage()

				if err != nil {
					log.Println(err.Error())
					return
				}

				messages <- msg
			}
		}()

		for message := range messages {
			go func(msg []byte) {
				time.Sleep(time.Second * 1)

				m := fiber.Map{
					"type":    "translation",
					"content": string(msg),
				}
				c.WriteJSON(m)
			}(message)

			go func(msg []byte) {
				time.Sleep(time.Second * 5)
				url := cloud.SignUrl("recording.mp3", 60*10)

				m := fiber.Map{
					"type":    "audio",
					"content": url,
				}

				c.WriteJSON(m)
			}(message)
		}
	})
}
