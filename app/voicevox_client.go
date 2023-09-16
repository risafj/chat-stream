package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type VoiceVoxClient struct {
	baseUrl string
	cid     string
}

func CreateVoiceVoxClient(baseUrl, cid string) *VoiceVoxClient {
	return &VoiceVoxClient{
		baseUrl: baseUrl,
		cid:     cid,
	}
}

func (c *VoiceVoxClient) GetAudio(msg string) (io.ReadCloser, error) {
	aq, err := c.getAudioQuery(msg)
	if err != nil {
		return nil, err
	}
	return c.getSynthesizedVoice(aq)
}

func (c *VoiceVoxClient) getAudioQuery(text string) (io.ReadCloser, error) {
	// Make an http request to BASE_URL + AUDIO_QUERY_ENDPOINT
	u, err := url.ParseRequestURI(c.baseUrl)
	if err != nil {
		return nil, err
	}
	u.Path = "/audio_query"
	u.RawQuery = url.Values{
		"speaker": {c.cid},
		"text":    {text},
	}.Encode()
	res, e := http.Post(u.String(), "application/json", nil)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != http.StatusOK {
		return nil, e
	}
	return res.Body, nil
}

func (c *VoiceVoxClient) getSynthesizedVoice(audioQuery io.ReadCloser) (io.ReadCloser, error) {
	u, err := url.ParseRequestURI(c.baseUrl)
	if err != nil {
		return nil, err
	}
	u.Path = "/synthesis"
	u.RawQuery = url.Values{
		"speaker": {c.cid},
	}.Encode()
	res, e := http.Post(u.String(), "application/json", audioQuery)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	return res.Body, nil
}
