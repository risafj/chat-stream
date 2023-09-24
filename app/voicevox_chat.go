package main

import (
	"log"
)

func voiceChat(voiceClient *VoiceVoxClient, chatClient *ChatGPTClient, audioPlayerClient *AudioPlayerClient) {
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

		err = audioPlayerClient.Play(audioClip)
		if err != nil {
			log.Fatalf("Error playing audio: %v", err)
		}
		log.Printf("Response: %s\n", output)
	}
}
