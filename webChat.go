package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// model: "gpt-3.5-turbo",
// messages: [{"role": "user", "content": "Say this is a test!"}],
// temperature: 0.7

type messageList struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// type postBody struct {
// 	Model       string   `json:"model"`
// 	Messages    []string `json:"messages"`
// 	Temperature float32  `json:"temperature"`
// }

type postBody struct {
	Model    string `json:"model"`
	Messages struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	}
	Temperature float32 `json:"temperature"`
}

type escritorWeb struct{}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))

	return len(p), nil
}

func isError(e error) {
	if e != nil {
		log.Fatal("Error-> ", e)
	}
}

func getChaT() {
	client := &http.Client{}
	// const apiKey = "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"
	// r, err := http.Get("https://api.openai.com/v1/models")
	req, err := http.NewRequest("GET", "https://api.openai.com/v1/models", nil)

	// r.Header.Add("Content-Type", "application/json")
	req.Header.Add("organization", "org-nMx2jYq1H2AffxOU1nAiZWy8")
	req.Header.Add("Authorization", "Bearer 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")

	isError(err)
	r, err := client.Do(req)
	e := escritorWeb{}

	io.Copy(e, r.Body)
	// fmt.Println(r.Body)
}

func postChaT() {

	// messagelist := []messageList{
	// 	{
	// 		Role:    "user",
	// 		Content: "Say this is a test!",
	// 	},
	// }
	// messagelis := messageList{
	// 	Role:    "user",
	// 	Content: "Say this is a test!",
	// }
	// marshalled1, err := json.Marshal(messagelis)
	// if err != nil {
	// 	log.Fatalf("impossible to marshall teacher: %s", err)
	// }
	// fmt.Println(string(marshalled1))

	postbody := postBody{
		Model: "gpt-3.5-turbo",
		// Messages: []string{string(marshalled1)},
		// Messages: []string{messagelis.Content, messagelis.Role},
		// Messages:    []string{array[0]},
		// Messages.Role: "",
		Messages: struct {
			Role    string "json:\"role\""
			Content string "json:\"content\""
		}{},
		Temperature: 0.7,
	}
	marshalled, err := json.Marshal(postbody)
	if err != nil {
		log.Fatalf("impossible to marshall teacher: %s", err)
	}
	fmt.Println(string(marshalled))

	client := &http.Client{}
	// const apiKey = "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"
	// r, err := http.Get("https://api.openai.com/v1/models")

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(marshalled))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("organization", "org-nMx2jYq1H2AffxOU1nAiZWy8")
	req.Header.Add("Authorization", "Bearer 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")

	isError(err)
	r, err := client.Do(req)
	e := escritorWeb{}

	io.Copy(e, r.Body)
	// fmt.Println(r.Body)
}

func main() {
	getChaT()
	// postChaT()
}
