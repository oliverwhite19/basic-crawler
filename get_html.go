package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {

	htmlResult, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("got Network error: %v", err)
	}

	if htmlResult.StatusCode >= 400 {
		return "", fmt.Errorf("got HTTP error: %s", htmlResult.Status)

	}

	if !strings.Contains(htmlResult.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", htmlResult.Header.Get("content-type"))
	}
	result, err := io.ReadAll(htmlResult.Body)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
