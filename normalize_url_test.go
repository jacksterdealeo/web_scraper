package main

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
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
		// Not using this test because the port could be useful. idk
		{
			name:     "removing the default port",
			inputURL: "http://example.com:80/",
			expected: "example.com",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
