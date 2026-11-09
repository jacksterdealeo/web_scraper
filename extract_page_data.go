package main

import (
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	Heading        string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) (d PageData) {
	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		log.Println(err)
	}
	d.URL = pageURL
	d.Heading = getHeadingFromHTML(html)
	d.FirstParagraph = getFirstParagraphFromHTML(html)
	d.OutgoingLinks, err = getURLsFromHTML(html, parsedURL)
	if err != nil {
		log.Println(err)
	}
	d.ImageURLs, err = getImagesFromHTML(html, parsedURL)
	if err != nil {
		log.Println(err)
	}
	return
}
