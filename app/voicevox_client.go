package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type voiceVoxClient struct {
	baseUrl string
	cid     string
}

func CreateVoiceVoxClient(baseUrl, cid string) *voiceVoxClient {
	return &voiceVoxClient{
		baseUrl: baseUrl,
		cid:     cid,
	}
}

func (c *voiceVoxClient) GetAudio(msg string) (io.ReadCloser, error) {
	aq, err := c.getAudioQuery(msg)
	if err != nil {
		return nil, err
	}
	return c.getSynthesizedVoice(aq)
}

func (c *voiceVoxClient) getAudioQuery(text string) (io.ReadCloser, error) {
	// Make an http request to BASE_URL + AUDIO_QUERY_ENDPOINT
	params := url.Values{}
	params.Add("speaker", c.cid)
	params.Add("text", text)
	u, _ := url.ParseRequestURI(c.baseUrl)
	u.Path = "/audio_query"
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	res, e := http.Post(urlStr, "application/json", nil)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != http.StatusOK {
		return nil, e
	}
	return res.Body, nil
}

func (c *voiceVoxClient) getSynthesizedVoice(audioQuery io.ReadCloser) (io.ReadCloser, error) {
	params := url.Values{}
	params.Add("speaker", c.cid)
	u, _ := url.ParseRequestURI(c.baseUrl)
	u.Path = "/synthesis"
	u.RawQuery = params.Encode()
	urlStr := fmt.Sprintf("%v", u)
	res, e := http.Post(urlStr, "application/json", audioQuery)
	if e != nil {
		return nil, e
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", res.StatusCode)
	}
	return res.Body, nil
}
