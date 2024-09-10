package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	res, err := getHtml(args[0])
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(res)

}
