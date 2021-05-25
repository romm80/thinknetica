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
	storage := "./storage.json"
	urls := []string{"https://golang.org/", "https://go.dev/"}
	gs := gosearch{}
	gs.scanner = spider.New()
	docs := []crawler.Document{}
	f, err := os.Open(storage)
	if err != nil {
		f, err = os.Create(storage)
		if err != nil {
			log.Fatal(err)
		}
		chRes, _ := gs.scanner.BatchScan(urls, 2, 2)
		for elem := range chRes {
			elem.ID = len(docs) + 1
			docs = append(docs, elem)
			err = add(f, elem)
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		docs, err = scan(f)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer f.Close()
	sort.Slice(docs, func(i, j int) bool { return docs[i].ID <= docs[j].ID })
	gs.index.Create(docs)

	var keyword = flag.String("s", "", "keyword")
	flag.Parse()
	*keyword = "why"
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

func add(w io.Writer, elem crawler.Document) error {
	b, err := json.Marshal(elem)
	if err != nil {
		return err
	}
	_, err = w.Write(append(b, '\n'))
	return err
}

func scan(r io.Reader) ([]crawler.Document, error) {
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
	return docs, scanner.Err()
}
