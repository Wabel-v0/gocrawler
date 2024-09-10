package main

import (
	"log"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawUrl string) []string {
	baseUrl, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatalf("Error parsing URL: %v", err)

	}

	htmlReader := strings.NewReader(htmlBody)
	doc, err := html.Parse(htmlReader)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}
	urls := []string{}
	var findUrls func(*html.Node)
	findUrls = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href, err := url.Parse(a.Val)
					if err != nil {
						log.Fatalf("Error parsing URL: %v", err)
						continue
					}
					// ResolveReference will resolve the relative url to the base url
					// i used to check if the url has a prefix of http or https , the resolveReference is much better :)
					baseUrl = baseUrl.ResolveReference(href)
					urls = append(urls, baseUrl.String())

				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			findUrls(c)
		}
	}
	findUrls(doc)
	return urls

}
