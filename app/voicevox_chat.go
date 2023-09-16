package main

import (
	"io"
	"log"
	"os"
)

func voiceChat(voiceClient *VoiceVoxClient, chatClient *ChatGPTClient) {
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
		err = outputFile.Close()
		if err != nil {
			log.Fatalf("Error closing the file: %f", err)
			return
		}
		log.Printf("Response: %s\n", output)
	}
}
