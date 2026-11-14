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

func (cfg *config) crawlPage(rawBaseURL, rawCurrentURL string) {
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

	if _, exists := cfg.pages[rawCurrentURLNormalized]; exists {
		return
	}
	pageText, err := getHTML(rawCurrentURL)
	if err != nil {
		crawlPageError("getHTML Failed. Will not bother recording details.", err)
		return
	}
	fmt.Printf("%v -> Searching %v\n", composedTime(), rawCurrentURLNormalized)
	gottenURLs, err := getURLsFromHTML(pageText, rawCurrentURLParsed)
	if err != nil {
		crawlPageError("", err)
	}
	gottenImages, err := getImagesFromHTML(pageText, rawCurrentURLParsed)
	if err != nil {
		crawlPageError("", err)
	}

	cfg.pages[rawCurrentURLNormalized] = PageData{
		URL:            rawCurrentURLNormalized,
		Heading:        getHeadingFromHTML(pageText),
		FirstParagraph: getFirstParagraphFromHTML(pageText),
		OutgoingLinks:  gottenURLs,
		ImageURLs:      gottenImages,
	}

	// Search the sites found.
	for _, url := range gottenURLs {
		cfg.crawlPage(rawBaseURL, url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (shouldVisit bool) {
	if _, exists := cfg.pages[normalizedURL]; exists {
		return false
	}
	return true
}
