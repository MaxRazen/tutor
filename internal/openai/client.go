package openai

import (
	"github.com/MaxRazen/tutor/internal/config"
	ai "github.com/sashabaranov/go-openai"
)

func NewClient() *ai.Client {
	return ai.NewClient(config.GetEnv(config.OPENAI_SECRET_KEY, ""))
}
