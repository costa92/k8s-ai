package ai

import "github.com/sashabaranov/go-openai"

type ChatMessages []openai.ChatCompletionMessage

var MessageStore ChatMessages

func (cm *ChatMessages) AddFor(role string, msg string) {
	*cm = append(*cm, openai.ChatCompletionMessage{
		Role:    role,
		Content: msg,
	})
}

func (cm *ChatMessages) ToMessage() []openai.ChatCompletionMessage {
	return append([]openai.ChatCompletionMessage(nil), *cm...)
}
