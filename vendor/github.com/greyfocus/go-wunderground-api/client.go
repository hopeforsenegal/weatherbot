package wunderground_api

import (
	"bytes"
	"errors"
	"net/http"
)

const (
	API_BASE_URL = "http://api.wunderground.com/api/"
)

type Client interface {
	runRequest(Request) Response
}

type JsonClient struct {
	ApiKey string
}

func (c JsonClient) buildRequestUrl(req *Request) string {
	buffer := bytes.NewBufferString(API_BASE_URL)
	buffer.WriteString(c.ApiKey)

	for _, f := range req.Features {
		buffer.WriteString("/")
		buffer.WriteString(f)
	}

	buffer.WriteString("/q/")
	buffer.WriteString(req.Location)

	buffer.WriteString(".json")

	return buffer.String()
}

func (c JsonClient) Execute(req *Request) (*Response, error) {
	url := c.buildRequestUrl(req)

	// fmt.Println("Request from " + url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("invalid HTTP status code: " + resp.Status)
	}

	defer resp.Body.Close()
	return parseWeatherResponse(resp.Body)
}
