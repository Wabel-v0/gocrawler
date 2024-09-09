package main

import (
	"log"
	"net/url"
)

func NormalizeUrl(rawUrl string) string {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	return parsedUrl.Host + parsedUrl.Path
}
