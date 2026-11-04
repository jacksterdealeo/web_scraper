package main

import (
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
