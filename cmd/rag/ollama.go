package main

import (
	"bytes"
	"io"
	"net/http"
)

// Generator struct to hold any necessary configuration.
type Generator struct {
	URL     string
	Headers map[string]string
}

// NewGenerator creates a new instance of Generator with default settings.
func NewGenerator() *Generator {
	return &Generator{
		URL: "http://localhost:11434/api/generate",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

func NewEmbedding() *Generator {
	return &Generator{
		URL: "http://localhost:11434/api/embedd",
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}

// Generate makes a POST request to the specified URL with the given body.
func (g *Generator) Generate(body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", g.URL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range g.Headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
