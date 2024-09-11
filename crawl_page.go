package main

import (
	"fmt"
	"log"
	"net/url"
)

func crawlPage(rawBaseUrl, rawCurrentUrl string, pages map[string]int) {
	baseUrl, err := url.Parse(rawBaseUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	currentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)
	}
	if baseUrl.Hostname() != currentUrl.Hostname() {
		fmt.Println("Not the same domain")
		fmt.Println(pages)
		return

	}
	normCurrentUrl := NormalizeUrl(rawCurrentUrl)

	if _, ok := pages[normCurrentUrl]; ok {
		pages[normCurrentUrl]++
		return
	} else {
		pages[normCurrentUrl] = 1
	}
	htmlResponseBody, err := getHtml(rawCurrentUrl)
	if err != nil {
		log.Fatalf("Error getting HTML: %v", err)

	}

	fmt.Printf("crawling %s\n", rawCurrentUrl)
	urls := getURLsFromHTML(htmlResponseBody, rawCurrentUrl)
	for _, u := range urls {
		crawlPage(rawBaseUrl, u, pages)
	}

}
