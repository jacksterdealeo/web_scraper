package main

import (
	"fmt"
	"net/url"
)

// Only for use in crawlPage. Prints out the errors it encounters.
func crawlPageError(msg string, e error) {
	if e == nil {
		return
	}
	fmt.Printf("%v -> crawlPage function had an error: %v\n", composedTime(), e)
	if msg != "" {
		fmt.Printf("  -> %v\n", msg)
	}
}

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	rawBaseURLParsed, err := url.Parse(rawBaseURL)
	if err != nil {
		crawlPageError("", err)
		return
	}
	rawCurrentURLParsed, err := url.Parse(rawCurrentURL)
	if err != nil {
		crawlPageError("", err)
		return
	}
	if rawBaseURLParsed.Hostname() != rawCurrentURLParsed.Hostname() {
		crawlPageError("", err)
		return
	}
	rawCurrentURLNormalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		crawlPageError("", err)
		return
	}
	_, exists := pages[rawCurrentURLNormalized]
	if exists {
		pages[rawCurrentURLNormalized] += 1
		return
	} else {
		pages[rawCurrentURLNormalized] = 1
	}
	page, err := getHTML(rawCurrentURL)
	if err != nil {
		crawlPageError("getHTML Failed.", err)
		return
	}
	fmt.Printf("%v -> Searching %v\n", composedTime(), rawCurrentURLNormalized)
	URLs, err := getURLsFromHTML(page, rawCurrentURLParsed)
	if err != nil {
		crawlPageError("", err)
		return
	}
	for _, url := range URLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
