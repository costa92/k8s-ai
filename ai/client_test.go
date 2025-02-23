// Copyright 2025 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/k8s-ai.

package ai

import (
	"context"
	"fmt"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestAiClient(t *testing.T) {
	client := NewDeepSeekAiClient()
	res, err := client.Client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: "deepseek-chat",
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleUser, Content: "Hello, world!"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Choices[0].Message.Content)
}

func TestChat(t *testing.T) {
	MessageStore.AddFor(openai.ChatMessageRoleSystem, "你是一个足球领域的专家，请尽可能地帮我回答与足球相关的问题")
	MessageStore.AddFor(openai.ChatMessageRoleUser, "C罗是哪个国家的足球运动员？")
	MessageStore.AddFor(openai.ChatMessageRoleAssistant, "C罗是葡萄牙足球运动员。")
	MessageStore.AddFor(openai.ChatMessageRoleUser, "内马尔呢？")
	msg := MessageStore.ToMessage()
	client := NewDeepSeekAiClient()
	res := client.Chat(msg)
	fmt.Println(res.Content)
}
