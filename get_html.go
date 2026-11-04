package main

import (
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
