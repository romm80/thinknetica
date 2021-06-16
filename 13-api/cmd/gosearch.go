package main

import (
	"log"
	"net/http"
	"thinknetica/13-api/pkg/api"
	"thinknetica/13-api/pkg/crawler"
	"thinknetica/13-api/pkg/crawler/spider"
	"thinknetica/13-api/pkg/storage"
)

type gosearch struct {
	scanner crawler.Interface
	api     *api.API
}

func main() {
	gs := gosearch{}
	gs.Init()
	log.Fatal(http.ListenAndServe(":8000", gs.api.Router()))
}

func (gs *gosearch) Init() {
	gs.scanner = spider.New()
	urls := []string{"https://golang.org/", "https://go.dev/"}
	chRes, _ := gs.scanner.BatchScan(urls, 2, 2)
	i := 1
	docs := []crawler.Document{}
	for elem := range chRes {
		elem.ID = i
		i++
		docs = append(docs, elem)
	}
	gs.api = api.New(storage.New(docs))
}
