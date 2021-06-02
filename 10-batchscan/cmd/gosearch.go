package main

import (
	"flag"
	"fmt"
	"sort"
	"thinknetica/10-batchscan/pkg/crawler"
	"thinknetica/10-batchscan/pkg/crawler/spider"
	"thinknetica/10-batchscan/pkg/index"
)

type gosearch struct {
	scanner crawler.Interface
	index   index.Index
}

func main() {
	urls := []string{"https://golang.org/", "https://go.dev/"}
	gs := gosearch{}
	gs.scanner = spider.New()
	docs := []crawler.Document{}

	chRes, _ := gs.scanner.MyBatchScan(urls, 2, 2)
	for elem := range chRes {
		elem.ID = len(docs) + 1
		docs = append(docs, elem)
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].ID <= docs[j].ID })
	gs.index.Create(docs)

	var keyword = flag.String("s", "", "keyword")
	flag.Parse()
	if *keyword != "" {
		res := gs.index.Search(*keyword)
		for _, id := range res {
			f := sort.Search(len(docs), func(i int) bool { return docs[i].ID >= id })
			if docs[f].ID == id {
				fmt.Printf("ID: %v, url: %v, title: %v\n", docs[f].ID, docs[f].URL, docs[f].Title)
			}
		}
	}
}
