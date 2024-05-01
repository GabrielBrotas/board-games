package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OpenAIRequest struct {
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

type OpenAIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func generateWordFromOpenAI(category string) (string, error) {
	openAIURL := "https://api.openai.com/v1/engines/davinci/completions"
	openAIKey := os.Getenv("OPENAI_KEY")

	if openAIKey == "" {
		return "", fmt.Errorf("OPENAI_KEY environment variable not set")
	}

	client := &http.Client{}
	prompt := "Please provide a word related to the theme: " + category
	if category == "" {
		prompt = "Please provide a random word"
	}

	data := OpenAIRequest{
		Prompt:    prompt,
		MaxTokens: 1,
	}

	body, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", openAIURL, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openAIKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response OpenAIResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) > 0 {
		return response.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no word returned from OpenAI")
}
