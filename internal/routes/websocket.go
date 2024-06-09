package routes

import (
	"log"
	"strconv"

	"github.com/MaxRazen/tutor/internal/room"
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

type incomingMessage struct {
	Type int
	Body []byte
}

func RoomWebsocketHandler() routeHandler {
	return websocket.New(func(c *websocket.Conn) {
		if c.Locals("allowed") == false {
			c.Close()
			return
		}

		roomId, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			log.Println("ws: roomId can not be parsed", roomId)
			// TODO: write msg to client
			c.Close()
			return
		}

		defer func() {
			println("connection per room", c.Params("id"), "is closed")
		}()

		messages := make(chan incomingMessage)
		go func() {
			defer close(messages)

			for {
				msgType, msg, err := c.ReadMessage()

				println("received message with type", msgType, len(msg))

				if err != nil {
					log.Println(err.Error())
					return
				}

				input := incomingMessage{
					Type: msgType,
					Body: msg,
				}

				messages <- input
			}
		}()

		msgWriter := func(data []byte) {
			c.WriteMessage(websocket.TextMessage, data)
		}

		go room.SendInitMessage(roomId, msgWriter)

		for message := range messages {
			if message.Type == websocket.BinaryMessage {
				// save recording asynchronously
				// TODO: add context
				go room.AcceptVoiceCommand(roomId, message.Body, msgWriter)
			}
		}
	})
}
