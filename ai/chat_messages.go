// Copyright 2025 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/k8s-ai.

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
