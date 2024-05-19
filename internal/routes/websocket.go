package routes

import (
	"log"
	"strconv"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/utils"
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

type IncomingMessage struct {
	Type int
	Body []byte
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

		messages := make(chan IncomingMessage)
		go func() {
			defer close(messages)

			for {
				msgType, msg, err := c.ReadMessage()

				println("received message with type", msgType, len(msg))

				if err != nil {
					log.Println(err.Error())
					return
				}

				input := IncomingMessage{
					Type: msgType,
					Body: msg,
				}

				messages <- input
			}
		}()

		for message := range messages {

			go func(msg IncomingMessage) {
				if msg.Type != websocket.BinaryMessage {
					return
				}

				filename := utils.GenerateRandomString(8) + ".ogg"

				err := cloud.Upload("rooms/1001/"+filename, msg.Body)

				if err != nil {
					println(err.Error())
				}

				m := fiber.Map{
					"type":    "translation",
					"content": "message processed with type: " + strconv.Itoa(msg.Type),
				}
				c.WriteJSON(m)
			}(message)

			// go func(msg []byte) {
			// 	time.Sleep(time.Second * 5)
			// 	url := cloud.SignUrl("recording.mp3", 60*10)

			// 	m := fiber.Map{
			// 		"type":    "audio",
			// 		"content": url,
			// 	}

			// 	c.WriteJSON(m)
			// }(message)
		}
	})
}
