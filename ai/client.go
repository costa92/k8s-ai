package ai

import (
	"context"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
)

type AiClient struct {
	Client *openai.Client
}

func NewDeepSeekAiClient() *AiClient {
	config := openai.DefaultConfig(os.Getenv("deepseek_api_key"))
	config.BaseURL = "https://api.deepseek.com"
	client := openai.NewClientWithConfig(config)
	return &AiClient{Client: client}
}

func (a *AiClient) Chat(message []openai.ChatCompletionMessage) openai.ChatCompletionMessage {
	rsp, err := a.Client.CreateChatCompletion(context.TODO(), openai.ChatCompletionRequest{
		Model:    "deepseek-chat",
		Messages: message,
	})
	if err != nil {
		log.Println(err)
		return openai.ChatCompletionMessage{}
	}

	return rsp.Choices[0].Message
}
