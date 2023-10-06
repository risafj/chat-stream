package main

import (
	"io"
	"time"

	"github.com/ebitengine/oto/v3"
)

type AudioPlayerClient struct {
	context *oto.Context
}

func CreateAudioPlayerClient() *AudioPlayerClient {
	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.
	op := &oto.NewContextOptions{
		// Tried the sample rates suggested in the docs (44100 or 48000) but those did not work
		SampleRate: 24000,
		// Number of channels (aka locations) to play sounds from. Either 1 or 2. 1 is mono sound, and 2 is stereo.
		ChannelCount: 1,
		// Format of the source. Wav uses 16-bit integers.
		Format: oto.FormatSignedInt16LE,
	}

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic(err)
	}

	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	return &AudioPlayerClient{
		context: otoCtx,
	}
}

func (apc *AudioPlayerClient) Play(audioClip io.Reader) error {
	player := apc.context.NewPlayer(audioClip)
	player.Play()

	// Wait for the sound to finish playing
	for player.IsPlaying() {
		time.Sleep(time.Millisecond)
	}

	err := player.Close()
	return err
}
