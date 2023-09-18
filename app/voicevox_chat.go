package main

import (
	"io"
	"log"
	"time"

	"github.com/ebitengine/oto/v3"
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

		err = playAudio(audioClip)
		if err != nil {
			log.Fatalf("Error playing audio: %v", err)
		}

		//outputFile, err := os.Create("output.wav")
		//if err != nil {
		//	log.Fatalf("Error creating the output file: %v", err)
		//	return
		//}
		//// Copy the response body to the output file
		//_, err = io.Copy(outputFile, audioClip)
		//if err != nil {
		//	log.Fatalf("Error copying the response body to the output file: %f", err)
		//	return
		//}
		//err = outputFile.Close()
		//if err != nil {
		//	log.Fatalf("Error closing the file: %f", err)
		//	return
		//}
		log.Printf("Response: %s\n", output)
	}
}

func playAudio(audioClip io.Reader) error {
	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.
	op := &oto.NewContextOptions{}

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	op.SampleRate = 24000

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 1

	// Format of the source. go-mp3's format is signed 16bit integers.
	// NOTE: wav also uses 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		return err
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer(audioClip)
	player.Play()

	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	// If you don't want the player/sound anymore simply close
	err = player.Close()
	return err
}
