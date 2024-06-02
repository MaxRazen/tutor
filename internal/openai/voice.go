package openai

import (
	"context"
	"io"

	"github.com/MaxRazen/tutor/internal/cloud"
	ai "github.com/sashabaranov/go-openai"
)

func Speech2Text(storObjName string, ctx context.Context) (ai.AudioResponse, error) {
	fileReader, err := cloud.GetObjectReader(ctx, storObjName)

	if err != nil {
		return ai.AudioResponse{}, err
	}

	c := NewClient()
	req := ai.AudioRequest{
		Model:    ai.Whisper1,
		FilePath: storObjName,
		Format:   ai.AudioResponseFormatSRT,
		Reader:   fileReader,
	}

	return c.CreateTranscription(ctx, req)
}

func Text2Speech(text string, ctx context.Context) ([]byte, error) {
	c := NewClient()
	req := ai.CreateSpeechRequest{
		Model:          ai.TTSModel1,
		Input:          text,
		Voice:          ai.VoiceAlloy,
		ResponseFormat: ai.SpeechResponseFormatOpus,
	}

	resp, err := c.CreateSpeech(ctx, req)

	if err != nil {
		return nil, err
	}

	defer resp.Close()

	return io.ReadAll(resp)
}
