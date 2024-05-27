package room

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/utils"
)

const (
	RecordingLinkTTL      = 3600
	ResponseTypeRecording = "recording"
	ResponseTypeError     = "error"
)

type response struct {
	Type          string `json:"type"`
	Timestamp     int64  `json:"ts"`
	Authorship    string `json:"authorship"`
	Content       string `json:"content"`
	Transcription string `json:"transcription"`
}

func AcceptVoiceCommand(roomId int, blob []byte, writer func(data []byte)) {
	path, err := saveVoiceRecording(roomId, blob)

	if err != nil {
		log.Printf("room/service: %v", err)
		writer(serverErrorResponse("your recording cannot be processed"))
		return
	}

	url := cloud.SignUrl(path, RecordingLinkTTL)

	r := response{
		Type:          ResponseTypeRecording,
		Timestamp:     time.Now().Unix(),
		Authorship:    "user",
		Content:       url,
		Transcription: "[cannot be recognized]",
	}
	data, err := json.Marshal(r)

	if err != nil {
		log.Printf("room/service: %v", err)
		writer(serverErrorResponse())
		return
	}

	writer(data)
}

func saveVoiceRecording(roomId int, blob []byte) (string, error) {
	filename := utils.GenerateRandomString(8) + ".ogg"
	path := fmt.Sprintf("rooms/%v/%v", roomId, filename)

	err := cloud.Upload(path, blob)

	if err != nil {
		return "", err
	}

	return path, nil
}

func serverErrorResponse(msg ...string) []byte {
	var m string = "Something went wrong"

	if len(msg) > 0 {
		m = msg[0]
	}

	r := response{
		Type:    ResponseTypeError,
		Content: m,
	}

	data, _ := json.Marshal(r)

	return data
}
