package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Message struct {
	Text string `json:"text"`
}

type Response struct {
	Text string `json:"text"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/chat", chatHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func chatHandler(w http.ResponseWriter, r *http.Request) {
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	query := strings.ReplaceAll(message.Text, " ", "+")
	url := fmt.Sprintf("https://api.openai.com/v1/engines/gpt-3.5-turbo/completions?prompt=%s&max_tokens=100", query)

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer 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")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(response)
}
