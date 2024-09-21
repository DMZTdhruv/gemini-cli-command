package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type GeminiRequest struct {
	Contents []Content `json:"contents"`
}

type Content struct {
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a question or prompt. Usage: gemini \"Your question here\"")
		os.Exit(1)
	}

	// Join all arguments into a single string for the prompt
	prompt := strings.Join(os.Args[1:], " ")

	// Replace this with your actual API key
	apiKey := ""
	if apiKey == "" {
		fmt.Println("Please set your Gemini API key in the code.")
		os.Exit(1)
	}

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-pro-latest:generateContent?key=" + apiKey

	requestBody := GeminiRequest{
		Contents: []Content{
			{
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
	}

	fmt.Println("Thinking...")
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Failed to create request body:", err)
		os.Exit(1)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Failed to make the API request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read the response:", err)
		os.Exit(1)
	}

	if resp.StatusCode != 200 {
		fmt.Printf("API request failed with status code %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var geminiResponse GeminiResponse
	if err := json.Unmarshal(body, &geminiResponse); err != nil {
		fmt.Println("Failed to parse the JSON response:", err)
		os.Exit(1)
	}

	if len(geminiResponse.Candidates) > 0 && len(geminiResponse.Candidates[0].Content.Parts) > 0 {
		response := geminiResponse.Candidates[0].Content.Parts[0].Text
		fmt.Printf("Gemini: %s\n", response)
	} else {
		fmt.Println("No response generated.")
	}
}
