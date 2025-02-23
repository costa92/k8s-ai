// Copyright 2025 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/k8s-ai.

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
