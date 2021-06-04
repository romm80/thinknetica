package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sort"
	"thinknetica/11-net/netsrv/pkg/crawler"
	"thinknetica/11-net/netsrv/pkg/crawler/spider"
	"thinknetica/11-net/netsrv/pkg/index"
)

type gosearch struct {
	scanner crawler.Interface
	index   index.Index
	docs    []crawler.Document
}

func main() {
	urls := []string{"https://golang.org/", "https://go.dev/"}
	gs := gosearch{}
	gs.scanner = spider.New()
	gs.docs = []crawler.Document{}

	chRes, _ := gs.scanner.BatchScan(urls, 2, 2)
	for elem := range chRes {
		elem.ID = len(gs.docs) + 1
		gs.docs = append(gs.docs, elem)
	}
	sort.Slice(gs.docs, func(i, j int) bool { return gs.docs[i].ID <= gs.docs[j].ID })
	gs.index.Create(gs.docs)

	listener, err := net.Listen("tcp4", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn, gs)
	}
}

func handler(conn net.Conn, gs gosearch) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, _, err := r.ReadLine()
		if err != nil {
			return
		}
		res := gs.index.Search(string(msg))
		for _, id := range res {
			f := sort.Search(len(gs.docs), func(i int) bool { return gs.docs[i].ID >= id })
			if gs.docs[f].ID == id {
				_, err = fmt.Fprintf(conn, "ID: %v, url: %v, title: %v\r\n", gs.docs[f].ID, gs.docs[f].URL, gs.docs[f].Title)
				if err != nil {
					return
				}
			}
		}
		_, err = conn.Write([]byte{'\r', '\n'})
		if err != nil {
			return
		}
	}
}
