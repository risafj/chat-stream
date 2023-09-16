package main

import (
	"flag"
	"io"
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

func main() {
	loadEnv()
	openaiAPIKey := os.Getenv(openAiApiKey)
	voiceVoxApiUrl := os.Getenv("VOICEVOX_API_URL")
	if openaiAPIKey == "" {
		log.Fatalf("%s not set", openAiApiKey)
	}
	flag.Parse()
	chatClient := CreateChatClient(openaiAPIKey, *isVoiceVox)
	if *isVoiceVox {
		voiceClient := CreateVoiceVoxClient(voiceVoxApiUrl, *characterVoiceId)
		for {
			input := getInputFromCommandLine()
			output, err := chatClient.SendMessage(input)
			if err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			audioClip, err := voiceClient.GetAudio(output)
			if err != nil {
				log.Fatalf("Error getting audio: %v", err)
			}
			outputFile, err := os.Create("output.wav")
			if err != nil {
				log.Fatalf("Error creating the output file: %v", err)
				return
			}
			// Copy the response body to the output file
			_, err = io.Copy(outputFile, audioClip)
			if err != nil {
				log.Fatalf("Error copying the response body to the output file: %f", err)
				return
			}
			outputFile.Close()
			log.Printf("Response: %s\n", output)
		}
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
