package main

import (
	"fmt"
	"sort"
)

type Page struct {
	Key   string
	Value int
}

func printReport(pages map[string]int, baseURL string) {
	var sorted []Page

	fmt.Println("========================================")
	fmt.Printf("Crawling report for %s\n", baseURL)
	fmt.Println("========================================")
	sorted = sortedPages(pages)
	for _, v := range sorted {
		fmt.Printf("Found %v internal links to %v \n", v.Value, v.Key)

	}

}

func sortedPages(pages map[string]int) []Page {
	pagesSlice := []Page{}
	for url, num := range pages {
		pagesSlice = append(pagesSlice, Page{Key: url, Value: num})
		sort.Slice(pagesSlice, func(i, j int) bool {
			if pagesSlice[i].Value == pagesSlice[j].Value {
				return pagesSlice[i].Key > pagesSlice[j].Key
			}
			return pagesSlice[i].Value > pagesSlice[j].Value

		})
	}
	return pagesSlice

}
