package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	if parsedURL.Host == "" {
		return "", fmt.Errorf("invalid URL")
	}

	combined := parsedURL.Host + parsedURL.Path
	normalized := strings.ToLower(combined)
	result := strings.TrimSuffix(normalized, "/")
	return result, nil
}
