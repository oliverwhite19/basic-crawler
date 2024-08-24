package main

import (
	"net/url"
	"strings"
)

func normalizeURL(initialUrl string) (actual string, err error) {
	result, err := url.Parse(initialUrl)
	if err != nil {
		return "", err
	}

	fullPath := result.Host + result.Path

	fullPath = strings.ToLower(fullPath)

	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}
