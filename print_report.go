package main

import (
	"fmt"
	"slices"
	"strings"
)

func printSpaces() {
	fmt.Println("=============================")
}

func getOrderedPages(pages map[string]int) map[int][]string {
	invertedPages := make(map[int][]string)

	for page, linkCount := range pages {
		_, ok := invertedPages[linkCount]
		if ok {
			invertedPages[linkCount] = append(invertedPages[linkCount], page)
		} else {
			invertedPages[linkCount] = []string{page}
		}
	}

	for key := range invertedPages {
		slices.SortFunc(invertedPages[key], func(a, b string) int {
			return strings.Compare(a, b)
		})
	}

	return invertedPages

}

func printReport(pages map[string]int, baseURL string) {

	printSpaces()
	fmt.Printf("REPORT for %s\n", baseURL)
	printSpaces()

	orderedPages := getOrderedPages(pages)

	keys := make([]int, 0, len(orderedPages))
	for k := range orderedPages {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	slices.Reverse(keys)

	for _, k := range keys {
		for _, page := range orderedPages[k] {
			fmt.Printf("Found %v internal links to %v\n", k, page)
		}
	}
}
