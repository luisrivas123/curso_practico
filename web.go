package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

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

func main() {
	r, err := http.Get("https://google.com")
	isError(err)
	e := escritorWeb{}

	io.Copy(e, r.Body)
	// fmt.Println(r.Body)
}
