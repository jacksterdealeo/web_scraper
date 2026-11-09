package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		// From Boot.dev:
		{
			name:     "Basic",
			html:     "<html><body><h1>Test Title</h1></body></html>",
			expected: "Test Title",
		},
		// My own:
		{
			name: "Get first <h1> tag.",
			html: `
<html>
  <h1>TEST 1</h1>
  <h2>TEST FAILED</h2>
  <body>
    <div>4</div>
    <div>9</div>
  </body>
</html>
`,
			expected: "TEST 1",
		},
		{
			name: "Get first <h2> tag.",
			html: `
<html>
  <h2>TEST 2</h2>
  <body>
    <div>1</div>
    <div>2</div>
  </body>
</html>
`,
			expected: "TEST 2",
		},
		{
			html: `
<html>
  <body>
    <div>1</div>
  </body>
</html>
`,
			expected: "",
		},
		{
			name:     "Invalid HTML.",
			html:     "Bananza>>>>><<><><",
			expected: "",
		},
	}

	for i, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			actual := getHeadingFromHTML(testcase.html)
			if actual != testcase.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, testcase.name, testcase.expected, actual)
				return
			}
		})
	}
}

func TestGetFirstParagraphFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		html     string
		expected string
	}{
		// From Boot.dev:
		{
			name: "Find main priority.",
			html: `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`,
			expected: "Main paragraph.",
		},
		// My own:
		{
			name: "No paragraph to find.",
			html: `<html><body>
		<main>
		</main>
	</body></html>`,
			expected: "",
		},
		{
			name: "Paragraph is outside of main.",
			html: `<html><body>
		<p>Outside paragraph.</p>
		<main>
		</main>
	</body></html>`,
			expected: "Outside paragraph.",
		},
	}
	for i, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			actual := getFirstParagraphFromHTML(testcase.html)
			if actual != testcase.expected {
				t.Errorf("Test %v - %s FAIL: expected: %v, actual: %v", i, testcase.name, testcase.expected, actual)
				return
			}
		})
	}
}

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		// From Boot.dev:
		{
			name:      "Absolute",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="https://crawler-test.com"><span>Boot.dev</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com"},
		},
		// From Me:
		{
			name:      "Absolute 2",
			inputURL:  "https://crawler-test.com/xyz.html",
			inputBody: `<html><body><a href="https://crawler-test.com/xyz.html"><span>FULL LINK TO XYZ PAGE</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/xyz.html"},
		},
		{
			name:      "Relative",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="/xyz.html"><span>XYZ PAGE</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/xyz.html"},
		},
		{
			name:      "Relative 2",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><a href="animal/xyz.html"><span>ANIMAL XYZ PAGE</span></a></body></html>`,
			expected:  []string{"https://crawler-test.com/animal/xyz.html"},
		},
		{
			name:      "Nothing",
			inputURL:  "https://crawler-test",
			inputBody: `<html><body><a href=""><span>ANIMAL XYZ PAGE</span></a></body></html>`,
			expected:  []string{},
		},
	}
	for i, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			baseURL, err := url.Parse(testcase.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v\n    This is a bad test.", i, testcase.name, err)
				return
			}
			actual, err := getURLsFromHTML(testcase.inputBody, baseURL)
			if err != nil && len(testcase.errorContains) == 0 {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, testcase.name, err)
				return
			} else if err != nil && !strings.Contains(err.Error(), testcase.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, testcase.name, testcase.errorContains)
				return
			} else if err == nil && testcase.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, testcase.name, testcase.errorContains)
				return
			}
			if !reflect.DeepEqual(actual, testcase.expected) {
				t.Errorf("expected %v, got %v", testcase.expected, actual)
			}
		})
	}
}

func TestGetImagesFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		// From Boot.dev:
		{
			name:      "Relative",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/logo.png" alt="Logo"></body></html>`,
			expected:  []string{"https://crawler-test.com/logo.png"},
		},
		// From Me:
		{
			name:      "Absolute",
			inputURL:  "https://crawler-test.com/xyz.html",
			inputBody: `<html><body><img src="https://crawler-test.com/xyz.png" alt="test image"></body></html>`,
			expected:  []string{"https://crawler-test.com/xyz.png"},
		},
		{
			name:      "Relative",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="/xyz.png" alt="test image"></body></html>`,
			expected:  []string{"https://crawler-test.com/xyz.png"},
		},
		{
			name:      "Relative 2",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body><img src="animal/xyz.png" alt="test image"></body></html>`,
			expected:  []string{"https://crawler-test.com/animal/xyz.png"},
		},
		{
			name:      "Nothing",
			inputURL:  "https://crawler-test.com",
			inputBody: `<html><body></body></html>`,
			expected:  []string{},
		},
	}
	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			baseURL, err := url.Parse(testcase.inputURL)
			if err != nil {
				t.Errorf("couldn't parse input URL: %v", err)
				return
			}
			actual, err := getImagesFromHTML(testcase.inputBody, baseURL)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !reflect.DeepEqual(actual, testcase.expected) {
				t.Errorf("expected %v, got %v", testcase.expected, actual)
			}
		})
	}
}
