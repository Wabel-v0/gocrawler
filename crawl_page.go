package main

import (
	"log"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentUrl string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if len(cfg.pages) > cfg.maxPages {
		return

	}
	currentUrl, err := url.Parse(rawCurrentUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)

	}
	if currentUrl.Hostname() != cfg.baseURL.Hostname() {
		return
	}
	normalizCurrentUrl := NormalizeUrl(rawCurrentUrl)
	isFirst := cfg.addPageVisit(normalizCurrentUrl)
	if !isFirst {
		return
	}
	htmlResponseBody, err := getHtml(rawCurrentUrl)
	if err != nil {
		log.Fatalf("Error getting HTML: %v", err)
	}

	urls := getURLsFromHTML(htmlResponseBody, rawCurrentUrl)
	for _, u := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)

	}

}
