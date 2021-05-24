package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"thinknetica/05-io/pkg/crawler"
	"thinknetica/05-io/pkg/crawler/spider"
	"thinknetica/05-io/pkg/index"
)

type gosearch struct {
	scanner crawler.Interface
	index   index.Index
}

func main() {
	file := ".\\storage.json"
	gs := gosearch{}
	gs.scanner = spider.New()
	docs := []crawler.Document{}
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		if err != nil {
			log.Fatalln(err)
		}
		defer f.Close()
		urls := []string{"https://golang.org/", "https://go.dev/"}
		chRes, _ := gs.scanner.BatchScan(urls, 2, 2)
		for elem := range chRes {
			elem.ID = len(docs) + 1
			docs = append(docs, elem)
			b, err := json.Marshal(elem)
			if err != nil {
				log.Fatal(err)
			}
			err = store(f, append(b, '\n'))
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		docs, err = get(f)
		if err != nil {
			log.Fatal(err)
		}
	}
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

func store(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}

func get(r io.Reader) ([]crawler.Document, error) {
	docs := []crawler.Document{}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var res crawler.Document
		err := json.Unmarshal(scanner.Bytes(), &res)
		if err != nil {
			return nil, err
		}
		docs = append(docs, res)
	}
	return docs, nil
}