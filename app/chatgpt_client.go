package main

import (
	"context"
	"log"

	"github.com/sashabaranov/go-openai"
)

type ChatGPTClient struct {
	Client              *openai.Client
	ctx                 context.Context
	maxTokensPerMessage int
	maxContext          int
	messages            []openai.ChatCompletionMessage
	isStreaming         bool
}

func CreateChatClient(apiKey string, isBlock bool) *ChatGPTClient {
	return &ChatGPTClient{
		Client:              openai.NewClient(apiKey),
		ctx:                 context.Background(),
		maxContext:          4096,
		maxTokensPerMessage: 500,
		isStreaming:         isBlock,
	}
}

func (c *ChatGPTClient) SendMessage(msg string) (string, error) {
	c.addMessageToMessages(msg, openai.ChatMessageRoleUser)
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 2000,
		Messages:  c.messages,
	}
	res, err := c.Client.CreateChatCompletion(c.ctx, req)
	if err != nil {
		log.Printf("ChatCompletion Error: %v\n", err)
		return "", err
	}
	message := res.Choices[0].Message.Content
	c.addMessageToMessages(message, openai.ChatMessageRoleAssistant)
	return message, nil
}

func (c *ChatGPTClient) addMessageToMessages(message string, role string) {
	// Add all existing tokens in message content
	var totalTokens int
	for _, msg := range c.messages {
		totalTokens += len(msg.Content)
	}
	totalTokens += len(message)

	// if totalTokens is greater than maxContext - maxTokensPerMessage
	// remove the first message
	for totalTokens > c.maxContext-c.maxTokensPerMessage {
		removedMessageTokenCount := len(c.messages[0].Content)
		c.messages = c.messages[1:]
		totalTokens -= removedMessageTokenCount
	}

	c.messages = append(c.messages, openai.ChatCompletionMessage{
		Role:    role,
		Content: message,
	})
}

func (c *ChatGPTClient) Stream(msg string) (*openai.ChatCompletionStream, error) {
	c.addMessageToMessages(msg, openai.ChatMessageRoleUser)

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: c.maxTokensPerMessage,
		Messages:  c.messages,
		Stream:    true,
	}
	return c.Client.CreateChatCompletionStream(c.ctx, req)
}
