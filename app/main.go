package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// flags
// if --Block, use block responses
var isBlockFormat = flag.Bool("block", false, "Use non-streaming block responses")

// Env vars
var openAiApiKey = "OPENAI_API_KEY"

func main() {
	loadEnv()
	openaiAPIKey := os.Getenv(openAiApiKey)
	if openaiAPIKey == "" {
		log.Fatalf("%s not set", openAiApiKey)
	}
	flag.Parse()
	client := CreateChatClient(openaiAPIKey, *isBlockFormat)
	if *isBlockFormat {
		for {
			input := getInputFromCommandLine()
			output, err := client.SendMessage(input)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			log.Printf("Response: %s\n", output)
		}
	} else {
		stream(client)
	}
}

func loadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
