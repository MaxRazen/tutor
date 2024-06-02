package room

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/openai"
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

type wsWriter func(data []byte)

func AcceptVoiceCommand(roomId int, audio []byte, writer wsWriter) {
	rh, err := processUserVoiceCommand(roomId, audio, writer)

	if err != nil {
		return
	}

	_, err = generateAssistantVoiceResponse(roomId, rh.Transcription, writer)

	if err != nil {
		return
	}

	log.Printf("User voice command accepted and processed successfully: room #%v", roomId)
}

func processUserVoiceCommand(roomId int, audio []byte, writer wsWriter) (rh RoomHistory, err error) {
	authorship := openai.RoleUser
	recordingPath, err := saveVoiceRecording(roomId, audio)

	if err != nil {
		log.Printf("room/service: saving record: %v", err)
		writer(serverErrorResponse("your recording cannot be processed"))
		return
	}

	recordingPublicUrl := cloud.SignUrl(recordingPath, RecordingLinkTTL)

	ctx := context.TODO()
	transcription := "[your recording cannot be recognized]"
	transcriptionResponse, err := openai.Speech2Text(recordingPath, ctx)

	if err != nil {
		log.Printf("room/service: speech2text: %v", err)
	} else {
		transcription = transcriptionResponse.Text
	}

	rh = RoomHistory{
		RoomId:        roomId,
		Authorship:    authorship,
		Transcription: transcription,
		Recording:     recordingPath,
	}
	if err = createRoomHistoryRecord(rh); err != nil {
		log.Printf("room/service: saving history[0]: %v", err)
	}

	r := response{
		Type:          ResponseTypeRecording,
		Timestamp:     time.Now().Unix(),
		Authorship:    authorship,
		Content:       recordingPublicUrl,
		Transcription: transcription,
	}
	data, err := json.Marshal(r)

	if err != nil {
		log.Printf("room/service: response marhaling[0]: %v", err)
		writer(serverErrorResponse())
		return
	}

	writer(data)

	return
}

func generateAssistantVoiceResponse(roomId int, command string, writer wsWriter) (rh RoomHistory, err error) {
	authorship := openai.RoleAssistant
	var chatMessages []openai.MessageContext

	log.Println("generateAssistantVoiceResponse text command tobe processed", command)

	chatMessages = append(chatMessages, openai.MessageContext{
		Role:    openai.RoleUser,
		Content: command,
	})

	ctx := context.TODO()
	chatResponses, err := openai.ChatCompletition(chatMessages, ctx)
	if err != nil {
		log.Printf("room/service: chat completition: %v", err)
		writer(serverErrorResponse())
		return
	}

	var concatenatedResponses string
	for _, cr := range chatResponses {
		concatenatedResponses += cr.Content
	}
	log.Println("chatGP4 response text", concatenatedResponses)
	answerAudio, err := openai.Text2Speech(concatenatedResponses, ctx)
	if err != nil {
		log.Printf("room/service: text2speech: %v", err)
		writer(serverErrorResponse())
		return
	}

	recordingPath, err := saveVoiceRecording(roomId, answerAudio)

	if err != nil {
		log.Printf("room/service: saving record: %v", err)
		writer(serverErrorResponse("your recording cannot be processed"))
		return
	}

	rh = RoomHistory{
		RoomId:        roomId,
		Authorship:    authorship,
		Transcription: concatenatedResponses,
		Recording:     recordingPath,
	}
	if err = createRoomHistoryRecord(rh); err != nil {
		log.Printf("room/service: saving history[1]: %v", err)
	}

	recordingPublicUrl := cloud.SignUrl(recordingPath, RecordingLinkTTL)

	r := response{
		Type:          ResponseTypeRecording,
		Timestamp:     time.Now().Unix(),
		Authorship:    authorship,
		Content:       recordingPublicUrl,
		Transcription: rh.Transcription,
	}
	data, err := json.Marshal(r)

	if err != nil {
		log.Printf("room/service: response marhaling[0]: %v", err)
		writer(serverErrorResponse())
		return
	}

	writer(data)

	return
}

func saveVoiceRecording(roomId int, audio []byte) (string, error) {
	filename := utils.GenerateRandomString(8) + ".ogg"
	path := fmt.Sprintf("rooms/%v/%v", roomId, filename)

	err := cloud.Upload(path, audio)

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
