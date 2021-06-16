package main

import (
	"log"
	"net/http"
	"sort"
	"thinknetica/12-http/pkg/crawler"
	"thinknetica/12-http/pkg/crawler/spider"
	"thinknetica/12-http/pkg/index"
	"thinknetica/12-http/pkg/webapp"
)

type gosearch struct {
	scanner crawler.Interface
	index   index.Index
	docs    []crawler.Document
}

func main() {
	gs := gosearch{}
	gs.Init()
	wa := webapp.Init(gs.index, gs.docs)
	log.Fatal(http.ListenAndServe(":8000", wa.Router))
}

func (g *gosearch) Init() {
	g.scanner = spider.New()
	g.docs = []crawler.Document{}
	urls := []string{"https://golang.org/", "https://go.dev/"}

	chRes, _ := g.scanner.BatchScan(urls, 2, 2)
	for elem := range chRes {
		elem.ID = len(g.docs) + 1
		g.docs = append(g.docs, elem)
	}
	sort.Slice(g.docs, func(i, j int) bool { return g.docs[i].ID <= g.docs[j].ID })
	g.index.Create(g.docs)
}
