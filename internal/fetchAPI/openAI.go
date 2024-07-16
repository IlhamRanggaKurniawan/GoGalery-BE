package fetchAPI

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type OpenAIResponse struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
}

type RequestBody struct {
	Model      string    `json:"model"`
	Messages   []Message `json:"messages"`
	Max_tokens uint      `json:"max_tokens"`
}

func FetchOpenAI(message string, responseChan chan OpenAIResponse, errorChan chan error) {

	requestBody := RequestBody{
		Model: "gpt-3.5-turbo-0125",
		Messages: []Message{
			{Role: "system", Content: "You are a helpful assistant"},
			{Role: "user", Content: message},
		},
		Max_tokens: 350,
	}

	jsonData, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))

	if err != nil {
		errorChan <- err
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ os.Getenv("OPENAI_KEY"))

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		errorChan <- err
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		errorChan <- err
		return
	}

	var openAIResponse OpenAIResponse

	err = json.Unmarshal(body, &openAIResponse)

	if err != nil {
		errorChan <- err
		return
	}

	responseChan <- openAIResponse
}