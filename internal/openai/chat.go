package openai

import (
	"context"
	"log"

	ai "github.com/sashabaranov/go-openai"
)

const (
	RoleUser      = ai.ChatMessageRoleUser
	RoleSystem    = ai.ChatMessageRoleSystem
	RoleAssistant = ai.ChatMessageRoleAssistant
)

type MessageContext = ai.ChatCompletionMessage

type MessageResponse struct {
	Role    string
	Content string
}

func ChatCompletition(messages []MessageContext, ctx context.Context) ([]MessageResponse, error) {
	c := NewClient()

	req := ai.ChatCompletionRequest{
		Model:    ai.GPT4o,
		Messages: messages,
	}

	res, err := c.CreateChatCompletion(ctx, req)

	if err != nil {
		log.Printf("openai/chat: %v", err)
		return nil, err
	}

	var responses []MessageResponse
	for _, m := range res.Choices {
		responses = append(responses, MessageResponse{
			Role:    m.Message.Role,
			Content: m.Message.Content,
		})
	}

	return responses, nil
}
