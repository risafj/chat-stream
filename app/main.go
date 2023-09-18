package main

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// flags
// if --vv, use block responses
var isVoiceVox = flag.Bool("vv", false, "Use voicevox models")
var characterVoiceId = flag.String("cid", "13", "Character voice id")

// Env vars
var openAiApiKey = "OPENAI_API_KEY"
var voiceVoxApiUrl = "VOICEVOX_API_URL"

func main() {
	loadEnv()
	openAiKey := os.Getenv(openAiApiKey)
	vvApiUrl := os.Getenv(voiceVoxApiUrl)
	if openAiKey == "" {
		log.Fatalf("%s not set", openAiApiKey)
	}
	flag.Parse()
	chatClient := CreateChatClient(openAiKey, *isVoiceVox)
	if *isVoiceVox {
		if vvApiUrl == "" {
			log.Fatalf("%s not set", voiceVoxApiUrl)
		}
		voiceClient := CreateVoiceVoxClient(vvApiUrl, *characterVoiceId)
		voiceChat(voiceClient, chatClient)
	} else {
		stream(chatClient)
	}
}

func loadEnv() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
