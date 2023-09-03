package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
	"io"
	"log"
	"os"
)

func stream() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(openaiAPIKey)
	ctx := context.Background()

	var messages []openai.ChatCompletionMessage
	for {
		var input string
		fmt.Printf("Message: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})

		req := openai.ChatCompletionRequest{
			Model:     openai.GPT3Dot5Turbo,
			MaxTokens: 2000,
			Messages:  messages,
			Stream:    true,
		}
		stream, err := client.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream Error: %v\n", err)
			break
		}

		fmt.Printf("Response: ")
		for {
			res, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nFin.")
				break
			}
			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				break
			}
			fmt.Printf("%v", res.Choices[0].Delta.Content)
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: res.Choices[0].Delta.Content,
			})
			continue
		}
		stream.Close()
	}
}
