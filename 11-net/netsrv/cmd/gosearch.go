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
	"time"
)

type gosearch struct {
	scanner crawler.Interface
	index   index.Index
	docs    []crawler.Document
}

func main() {
	gs := gosearch{}
	gs.Init()

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

func handler(conn net.Conn, gs gosearch) {
	conn.SetDeadline(time.Now().Add(time.Second * 10))
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
		conn.SetDeadline(time.Now().Add(time.Second * 10))
	}
}
