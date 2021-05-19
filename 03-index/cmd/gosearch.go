package main

import (
	"flag"
	"fmt"
	"sort"
	"thinknetica/03-index/pkg/crawler"
	"thinknetica/03-index/pkg/crawler/spider"
	"thinknetica/03-index/pkg/index"
)

type gosearch struct {
	scanner crawler.Interface
}

func main() {
	gs := gosearch{}
	gs.scanner = spider.New()
	urls := []string{"https://golang.org/", "https://go.dev/"}
	chRes, _ := gs.scanner.BatchScan(urls, 2, 2)
	docs := []crawler.Document{}
	idx := index.Index{}
	idx.New()
	for elem := range chRes {
		elem.ID = len(docs)
		idx.Add(elem.Title, elem.ID)
		docs = append(docs, elem)
	}
	sort.Slice(docs, func(i, j int) bool {return docs[i].ID <= docs[j].ID})

	var keyword = flag.String("s","","keyword")
	flag.Parse()
	if *keyword != "" {
		res := idx.Search(*keyword)
		for _, id := range res {
			f := sort.Search(len(docs), func(i int) bool {return docs[i].ID >= id})
			if docs[f].ID == id {
				fmt.Printf("ID: %v, url: %v, title: %v\n", docs[f].ID, docs[f].URL, docs[f].Title)
			}
		}
	}
}
