package main

import "testing"

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		expected  string
		expectErr bool
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove http scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "convert to lowercase",
			inputURL: "https://BLOG.BOOT.DEV/PATH",
			expected: "blog.boot.dev/path",
		},
		{
			name:      "invalid url",
			inputURL:  "not-a-url",
			expected:  "", // expect empty string when there's an error
			expectErr: true,
		},
		{
			name:     "root domain no trailing slash",
			inputURL: "https://blog.boot.dev",
			expected: "blog.boot.dev",
		},
		{
			name:     "capitals with trailing slash",
			inputURL: "HTTP://BLOG.BOOT.DEV/PATH/",
			expected: "blog.boot.dev/path",
		}, // add more test cases here
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeURL(tt.inputURL)
			if (err != nil && !tt.expectErr) || (err == nil && tt.expectErr) {
				t.Errorf("Test %v - %s FAIL: error expectation mismatch. Got error: %v, Expected error: %v", tt.name, tt.name, err, tt.expectErr)
				return
			}
			if got != tt.expected {
				t.Errorf("normalizeURL(%q) = %q, want %q", tt.inputURL, got, tt.expected)
			}
		})
	}
}
