package main

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"github.com/sashabaranov/go-openai"
// 	"io"
// 	"net/http"
// 	"os"
// 	"time"
// )

// func block(client *chatGPTClient) {
// 	ctx := context.Background()

// 	var messages []openai.ChatCompletionMessage
// 	for {
// 		input := getInputFromCommandLine()
// 		start := time.Now()

// 		messages = append(messages, openai.ChatCompletionMessage{
// 			Role:    openai.ChatMessageRoleUser,
// 			Content: input,
// 		})

// 		req := openai.ChatCompletionRequest{
// 			Model:     openai.GPT3Dot5Turbo,
// 			MaxTokens: 2000,
// 			Messages:  messages,
// 		}
// 		res, err := client.CreateChatCompletion(ctx, req)
// 		if err != nil {
// 			fmt.Printf("ChatCompletion Error: %v\n", err)
// 			break
// 		}

// 		response := res.Choices[0].Message.Content
// 		fmt.Printf("Response: %s\n", response)

// 		messages = append(messages, openai.ChatCompletionMessage{
// 			Role:    openai.ChatMessageRoleAssistant,
// 			Content: response,
// 		})

// 		requestBody, err := json.Marshal(VoiceRequest{VoiceLines: response})
// 		if err != nil {
// 			fmt.Printf("Json Error: %v\n", err)
// 			break
// 		}

// 		voiceResponse, err := http.Post("https://walrus-fluent-jaybird.ngrok-free.app/get-voice", "application/json", bytes.NewBuffer(requestBody))
// 		if err != nil {
// 			fmt.Printf("Voice Error: %v\n", err)
// 			break
// 		}
// 		defer voiceResponse.Body.Close()

// 		outputFile, err := os.Create("chat.wav")
// 		if err != nil {
// 			fmt.Println("Error creating the output file:", err)
// 			return
// 		}
// 		defer outputFile.Close()

// 		// Copy the response body to the output file
// 		_, err = io.Copy(outputFile, voiceResponse.Body)
// 		if err != nil {
// 			fmt.Println("Error copying the response body to the output file:", err)
// 			return
// 		}
// 		fmt.Println(time.Since(start))
// 	}
// }

// type VoiceRequest struct {
// 	VoiceLines string `json:"voice_lines"`
// }
