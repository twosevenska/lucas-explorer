package main

import (
	"flag"
	"fmt"
	"log"

	"lucas/crawl"
	"lucas/validate"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Please provide a url value.")
	}

	address := args[0]
	if !validate.URLAddress(address) {
		log.Fatal("Please provide a valid url value.")
	}

	page, err := crawl.Download(address)
	if err != nil {
		log.Fatalf("Failed to download page: %s", err.Error())
	}

	nLines := flag.Int("lines", 10, "Size of chunks to split from the page")

	for u, c := range crawl.ExtractURLS(page, *nLines) {
		fmt.Printf("%s %d\n", u, c)
	}
}
