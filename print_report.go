package main

import (
	"fmt"
	"sort"
)

type KeyValue struct {
	Key   string
	Value int
}

func printReport(pages map[string]int, baseURL string) {
	var sortedPages []KeyValue
	for k, v := range pages {
		sortedPages = append(sortedPages, KeyValue{k, v})
	}
	sort.SliceStable(sortedPages, func(i, j int) bool {
		return sortedPages[i].Value > sortedPages[j].Value
	})
	fmt.Println("========================================")
	fmt.Printf("Crawling report for %s\n", baseURL)
	fmt.Println("========================================")
	for _, v := range sortedPages {
		fmt.Printf("Found %v internal links to %v \n", v.Value, v.Key)

	}
}
