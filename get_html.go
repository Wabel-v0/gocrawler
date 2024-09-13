package main

import (
	"io"
	"log"
	"net/http"
)

func getHtml(rawUrl string) (string, error) {
	res, err := http.Get(rawUrl)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 399 {
		log.Fatalf("Error: %v", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)

	}
	return string(body), nil
}
