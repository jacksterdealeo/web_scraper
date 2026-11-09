package main

import (
	"fmt"
	"net/url"
	"slices"
	"strings"
)

// the normalized URL should not be used to make requests,
// it's just for comparing URLs to see if they are the same page.
func normalizeURL(urlString string) (string, error) {
	address, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	var result strings.Builder

	if address.User != nil {
		result.WriteString(address.User.String())
		result.WriteRune('@')
	}

	result.WriteString(strings.ToLower(address.Host))

	if len(address.Path) > 0 {
		var path []string
		if address.Path[len(address.Path)-1] == '/' {
			path = strings.Split(address.Path[:len(address.Path)-1], "/")
		} else {
			path = strings.Split(address.Path, "/")
		}
		for i := 0; i < len(path); {
			if c := strings.Count(path[i], "."); c > 0 {
				path = slices.Delete(path, i-(c-1), i+1)
				i -= (c - 1)
			} else {
				i += 1
			}
		}
		result.WriteString(strings.Join(path, "/"))
	}

	return result.String(), nil
}
