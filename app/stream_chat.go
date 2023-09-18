package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/sashabaranov/go-openai"
)

func stream(client *ChatGPTClient) {
	for {
		input := getInputFromCommandLine()
		stream, err := client.Stream(input)
		if err != nil {
			fmt.Printf("ChatCompletionStream Error: %v\n", err)
			break
		}
		fmt.Printf("Response: ")
		aggregatedStreamResponseMessage := ""
		for {
			res, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				fmt.Println("\nFin.")
				client.addMessageToMessages(aggregatedStreamResponseMessage, openai.ChatMessageRoleAssistant)
				break
			}
			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				break
			}
			fmt.Printf("%v", res.Choices[0].Delta.Content)
			aggregatedStreamResponseMessage += res.Choices[0].Delta.Content
			continue
		}
		stream.Close()
	}
}
