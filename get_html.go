package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "BootCrawler/1.0")

	client := http.Client{
		Transport:     http.DefaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	resp, err := client.Transport.RoundTrip(req)
	if err != nil {
		return "", fmt.Errorf("RoundTrip Failed. %w", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("Error from http status code: %v", resp.Status)
	}
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Error. Wrong Content-Type. Expected 'text/html', Got '%v'", resp.Header.Get("Content-Type"))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
