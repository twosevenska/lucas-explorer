package crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Download fetches a webpage and returns it as a string
func Download(address string) (string, error) {
	resp, err := http.Get(address)
	if err != nil {
		return "", fmt.Errorf("failed to get webpage: %s", err.Error())
	}
	defer resp.Body.Close()

	// Warning: This will not handle large webpages properly
	// https://groups.google.com/forum/#!topic/golang-nuts/sAwDldpkMGQ
	// It might be better to temporarily download to a file and scan it
	// line by line
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read webpage: %s", err.Error())

	}

	return string(html), nil
}

// ExtractURLS finds and extracts urls in a page from hyperlink tags
func ExtractURLS(page string, step int) map[string]int {

	lines := strings.Split(page, "\n")
	var chunks [][]string
	for step < len(lines) {
		lines, chunks = lines[step:], append(chunks, lines[0:step:step])
	}
	chunks = append(chunks, lines)

	urls := make(chan string, len(chunks))
	done := make(chan bool, len(chunks))
	for _, c := range chunks {
		go extractFromSection(c, urls, done)
	}

	addresses := make(map[string]int)
	dc := 0
	for {
		select {
		case u, more := <-urls:
			if !more {
				return addresses
			}
			addresses[u] += 1
		case <-done:
			dc += 1
			if dc == len(chunks) {
				close(urls)
			}
		}
	}
}

// extractFromSection goes through a section and tries to find any urls
func extractFromSection(chunk []string, urls chan<- string, done chan<- bool) {
	for _, l := range chunk {
		u, extracted := extractURL(l)
		if extracted {
			urls <- u
		}
	}
	done <- true
}

// extractURL attemps to find a url tag in a line and extract the address
func extractURL(line string) (string, bool) {
	startIndex := strings.Index(strings.ToLower(line), "href=\"")
	if startIndex == -1 {
		return "", false
	}

	trimmedPrefix := line[startIndex+6:]

	endIndex := strings.Index(trimmedPrefix, "\"")
	if endIndex == -1 {
		return "", false
	}

	result := trimmedPrefix[:endIndex]
	if result == "" {
		return "", false
	}
	return result, true
}
