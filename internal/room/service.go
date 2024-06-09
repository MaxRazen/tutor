package room

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/MaxRazen/tutor/internal/cloud"
	"github.com/MaxRazen/tutor/internal/openai"
	"github.com/MaxRazen/tutor/internal/utils"
	"github.com/MaxRazen/tutor/pkg/memcache"
)

const (
	RecordingLinkTTL      = 3600
	ResponseTypeRecording = "recording"
	ResponseTypeError     = "error"
	RecordingCachePrefix  = "audio:"
)

type response struct {
	Type          string `json:"type"`
	Timestamp     int64  `json:"ts"`
	Authorship    string `json:"authorship"`
	Content       string `json:"content"`
	Transcription string `json:"transcription"`
}

type wsWriter func(data []byte)

func GetPublicAudioLink(audio string) string {
	link, ok := memcache.Get(RecordingCachePrefix + audio)

	if !ok {
		link := cloud.SignUrl(audio, RecordingLinkTTL)
		key := RecordingCachePrefix + audio

		memcache.Set(key, link, time.Second*RecordingLinkTTL)
		return link
	}

	return link
}

func SendInitMessage(roomId int, writer wsWriter) {
	roomHistory, err := LoadRoomHistory(roomId)

	if err != nil {
		log.Println("room/service: initMsg:", err)
		return
	}

	// deliver room history if it present
	for _, rh := range roomHistory {
		if rh.Recording != "" {
			rh.Recording = GetPublicAudioLink(rh.Recording)
		}
		respMsg := response{
			Type:          ResponseTypeRecording,
			Timestamp:     rh.CreatedAt.Unix(),
			Authorship:    rh.Authorship,
			Content:       rh.Recording,
			Transcription: rh.Transcription,
		}
		data, err := json.Marshal(respMsg)
		if err != nil {
			log.Println("room/service:", err)
			continue
		}
		writer(data)
	}

	if len(roomHistory) > 0 && roomHistory[len(roomHistory)-1].Authorship == openai.RoleAssistant {
		// TODO: send generic question to continue the conversation
		// should not be written to history
		return
	}

	// init chatGPT request with context
	// TODO
}

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
		transcription = extractTranscription(transcriptionResponse.Text)
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

	callCtx := `You are an English tutor. Your name is Scarlett. You are between 25-30 years, you were born in London.
	This conversation is going to be performed in the way of a call
	(all the responses are provided by User and Assistant roles will be converted from audio to text and back).
	Keep your responses short enough and simulate real conversation. Try to use open questions to develop the dialog.
	If the User has made a mistake, correct him.`

	chatMessages = append(chatMessages, openai.MessageContext{
		Role:    openai.RoleSystem,
		Content: callCtx,
	})
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

// Extracting text from Whisper's response
// 1\n
// 00:00:00,000 --> 00:00:10,000\n
// text...\n
func extractTranscription(rawText string) string {
	ss := strings.Split(strings.Trim(rawText, "\n\t "), "\n")
	return strings.Join(ss[2:], " ")
}
