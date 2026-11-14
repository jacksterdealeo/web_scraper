package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	f := doc.FindMatcher(goquery.Single("h1")).Text()
	if f == "" {
		f = doc.FindMatcher(goquery.Single("h2")).Text()
	}
	return f
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	f := (doc.FindMatcher(goquery.Single("main")).FindMatcher(goquery.Single("p"))).Text()
	if f == "" {
		f = doc.FindMatcher(goquery.Single("p")).Text()
	}
	return f
}

// returns an un-normalized list of all the URLs found within the HTML
// on error, returns an empty slice
func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("couldn't read htmlBody: %w", err)
	}
	result := []string{}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		linkStr, found := s.Attr("href")
		if found && len(linkStr) > 0 {
			link, err := url.Parse(linkStr)
			if err != nil {
				return
			}
			if len(link.Host) == 0 {
				link.Host = baseURL.Host
			}
			if len(link.Scheme) == 0 {
				link.Scheme = baseURL.Scheme
			}
			result = append(result, link.String())
		}
	})
	return result, nil
}

// returns an un-normalized list of all the image URLs found within the HTML
// on error, returns an empty slice
func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, fmt.Errorf("couldn't read htmlBody: %w", err)
	}
	result := []string{}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		linkStr, found := s.Attr("src")
		if found && len(linkStr) > 0 {
			link, err := url.Parse(linkStr)
			if err != nil {
				return
			}
			if len(link.Host) == 0 {
				link.Host = baseURL.Host
			}
			if len(link.Scheme) == 0 {
				link.Scheme = baseURL.Scheme
			}
			result = append(result, link.String())
		}
	})
	return result, nil
}
