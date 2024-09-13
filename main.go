package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)

	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)

	} else if len(args) == 3 {
		fmt.Printf("starting crawl of: %v\n", args[0])
	}
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("invalid max concurrency")
		os.Exit(1)

	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("invalid max pages")
		os.Exit(1)

	}
	rawBaseURL := args[0]

	cfg := configer(rawBaseURL, maxConcurrency, maxPages)

	cfg.wg.Add(1)
	go cfg.crawlPage(args[0])
	cfg.wg.Wait()
	printReport(cfg.pages, rawBaseURL)

}
