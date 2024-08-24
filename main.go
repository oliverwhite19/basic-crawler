package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}

	if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseUrl := argsWithoutProg[0]
	maxConcurrency, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Println("maximum concurrency requires an integer")
		os.Exit(1)
	}

	maxPages, err := strconv.Atoi(argsWithoutProg[2])
	if err != nil {
		fmt.Println("maximum pages requires an integer")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s\n", baseUrl)
	cfg, err := configurePageCrawl(baseUrl, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseUrl)
	cfg.wg.Wait()

	printReport(cfg.pages, baseUrl)

}
