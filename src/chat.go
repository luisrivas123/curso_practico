package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiKey = "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"

func SendRequest(prompt string) string {
	url := "https://api.openai.com/v1/chat/completions"
	data := fmt.Sprintf(`{model: "gpt-3.5-turbo", "Messages":[{"role": "user", "content": "%s"}], "max_tokens": 100}`, prompt)
	fmt.Println(data)
	fmt.Println(prompt)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	// return result["choices"].([]map[string]interface{})[0]["text"].(string)
	fmt.Println(string(body))
	return client.Timeout.String()
}

func main() {
	response := SendRequest("¡Hola! ¿Cómo estás?")
	fmt.Println(response)
}
