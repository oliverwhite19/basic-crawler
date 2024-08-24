package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromNodeTree(node *html.Node) []string {
	resultingURLs := make([]string, 0)

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		resultingURLs = append(resultingURLs, getURLsFromNodeTree(child)...)
	}

	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attribute := range node.Attr {
			if attribute.Key == "href" {
				return append(resultingURLs, attribute.Val)
			}
		}

	}

	return resultingURLs
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	nodeTree, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	resultingURLs := getURLsFromNodeTree(nodeTree)

	fullURLs := make([]string, 0)
	for _, result := range resultingURLs {
		u, err := url.Parse(result)
		if err != nil {
			return nil, err
		}
		resolvedURL := baseURL.ResolveReference(u)
		fullURLs = append(fullURLs, resolvedURL.String())
	}

	return fullURLs, nil
}
