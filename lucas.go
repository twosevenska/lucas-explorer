package main

import (
	"flag"
	"fmt"
	"log"

	"lucas-explorer/crawl"
	"lucas-explorer/validate"
)

func main() {
	maxDepth := flag.Int("depth", 0, "How many time we follow each url in the same domain")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		log.Fatal("Please provide a url value.")
	}
	address := args[0]
	if !validate.URLAddress(address) {
		log.Fatal("Please provide a valid url value.")
	}

	for u := range processPage(address, 0, *maxDepth) {
		fmt.Printf("%s\n", u)
	}
}

func processPage(address string, curDepth, maxDepth int) map[string]bool {
	//fmt.Printf("Currently checking %s at depth: %d\n", address, curDepth)
	result := make(map[string]bool)

	if curDepth > maxDepth {
		return result
	}

	page, err := crawl.Download(address)
	if err != nil {
		log.Fatalf("Failed to download page: %s", err.Error())
	}

	result = crawl.ExtractURLS(page)
	for u := range result {
		for k, v := range processPage(u, curDepth+1, maxDepth) {
			result[k] = v
		}
	}

	return result
}
