package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)

	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)

	} else if len(args) == 1 {
		fmt.Printf("starting crawl of: %v\n", args[0])
	}
	cfg := configer(args[0], 5)

	cfg.wg.Add(1)
	go cfg.crawlPage(args[0])
	cfg.wg.Wait()
	for url, page := range cfg.pages {
		fmt.Printf("URL: %s, visits: %d\n", url, page)
	}

}
