package main

import (
	"flag"
	"fmt"
	"strings"
	"thinknetica/02-crawler/pkg/crawler"
	"thinknetica/02-crawler/pkg/crawler/spider"
)

type gosearch struct {
	scanner crawler.Interface
}

func main() {
	gs := gosearch{}
	gs.scanner = spider.New()
	urls := []string{"https://golang.org/", "https://go.dev/"}
	chRes, _ := gs.scanner.BatchScan(urls, 2, 1)
	docs := []crawler.Document{}
	for elem := range chRes {
		docs = append(docs, elem)
	}

	var keyword string
	flag.StringVar(&keyword, "s", "", "Слово для поиска")
	flag.Parse()
	if keyword != "" {
		for _, elem := range docs {
			if strings.Contains(elem.Title, keyword) {
				fmt.Println(elem.URL)
				fmt.Println(elem.Title)
			}
		}
	}
}
