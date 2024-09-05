package fetchAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int      `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *string  `json:"logprobs"`
	FinishReason string   `json:"finish_reason"`
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
	Model      string     `json:"model"`
	Messages   []Message `json:"messages"`
	Max_tokens uint64       `json:"max_tokens"`
}

func FetchOpenAI(prompt []Message) (OpenAIResponse, error) {

	requestBody := RequestBody{
		Model:    "gpt-4o-mini-2024-07-18",
		Messages: prompt,
		Max_tokens: 1000,
	}

	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		fmt.Println(err)
		return OpenAIResponse{}, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
		return OpenAIResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+os.Getenv("OPENAI_KEY"))

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return OpenAIResponse{}, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println(err)
		return OpenAIResponse{}, err
	}

	var openAIResponse OpenAIResponse

	err = json.Unmarshal(body, &openAIResponse)

	if err != nil {
		fmt.Println(err)
		return OpenAIResponse{}, err
	}

	return openAIResponse, nil
}
