# Goal

Use a combination of ChatGPT and VoiceVox to make cool chat robots.

## Setup

1) Install [VoiceVox Engine](https://github.com/VOICEVOX/voicevox_engine). I recommend using the pre-made docker images.
2) (If using VoiceVox for Japanese voices): Make the VoiceVox Engine available on port 50021. If you're using docker, you can do this with `docker run -p 50021:50021 -it [DOCKER CONTAINER YOU DOWNLOADED]`.
3) Sign up for the chatgpt api, and save your API key.

## Usage

Before beginning, we just need to set two environment variables inside a `.env` file.
```
OPENAI_API_KEY=[Your Api Key Here]
VOICEVOX_API_URL=[Your VoiceVox Endpoint Here]
```

If you don't plan to use VoiceVox and just want a streaming chat gpt service, you can leave the `VOICEVOX_API_URL` variable unset.

When run without any arguments, chat-stream will work approximately as ChatGPT does.
It will take messages from the command line and return chatGPT responses in a streaming mode.

Commands:
```
    --vv            Use VoiceVox to read responses out loud using a preset voice
    --vv-api        Run the program as an api server in VoiceVox mode. See Api Doc below.
    --cid           [Character ID] Use a specific VoiceVox character
```

An informal list of supported characters and corresponding ids can be found [here](https://puarts.com/?pid=1830).

Audio samples and licensing for those characters can be found [here](https://voicevox.hiroshiba.jp/).

## API Mode:

When run with the `--Api` flag, chat-stream will run as an api server.

The server has one endpoint.
 
| Method | Endpoint   | Header           | Body                  | Response  |
|--------|------------|------------------|-----------------------|-----------|
| Post   | /get-voice | contentType-json | {voice_lines: string} | audio/wav |
