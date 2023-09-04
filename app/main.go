package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// flags
// if --Block, use block responses
var isBlockFormat = flag.Bool("Block", false, "Use non-streaming block responses")

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
		// input := getInputFromCommandLine()
		panic("Not implemented")
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
