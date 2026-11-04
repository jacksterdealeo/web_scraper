package main

import (
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		// From Bootdev:
		{
			name:     "remove scheme",
			inputURL: "https://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		// From Wikipedia (https://en.wikipedia.org/wiki/URI_normalization) (with some changes):
		{
			name:     "converting the host to lowercase",
			inputURL: "HTTP://User@Example.COM/Foo",
			expected: "User@example.com/Foo",
		},
		{
			name:     "decoding percent-encoded triplets of unreserved characters",
			inputURL: "http://example.com/%7Efoo",
			expected: "example.com/~foo",
		},
		{
			name:     "removing dot-segments",
			inputURL: "http://example.com/foo/./bar/baz/../qux",
			expected: "example.com/foo/bar/qux",
		},
		{
			name:     "converting a \"/\" to an empty path",
			inputURL: "http://example.com/",
			expected: "example.com",
		},
		// From Boot.dev CH1 L3 solution: added check that error is caught
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidAddress`,
			expected:      "",
			errorContains: "couldn't parse URL",
		},
	}

	for i, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			actual, err := normalizeURL(testcase.inputURL)
			if err != nil && !strings.Contains(err.Error(), testcase.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, testcase.name, err)
				return
			} else if err != nil && testcase.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, testcase.name, err)
				return
			} else if err == nil && testcase.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, testcase.name, testcase.errorContains)
				return
			}
			if actual != testcase.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, testcase.name, testcase.expected, actual)
				return
			}
		})
	}
}
