package db

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"thinknetica/13-api/pkg/crawler"
)

type DB struct {
	docs []crawler.Document
	mu   sync.Mutex
}

func New(docs []crawler.Document) *DB {
	return &DB{docs: docs}
}

func (d *DB) All() *[]crawler.Document {
	return &d.docs
}

func (d *DB) Search(ids []int) []crawler.Document {
	res := []crawler.Document{}
	for _, id := range ids {
		i := sort.Search(len(d.docs), func(i int) bool { return d.docs[i].ID >= id })
		if d.docs[i].ID != id {
			continue
		}
		res = append(res, d.docs[i])
	}
	return res
}

func (d *DB) Add(doc crawler.Document) {
	d.mu.Lock()
	switch len(d.docs) {
	case 0:
		doc.ID = 1
	default:
		doc.ID = d.docs[len(d.docs)-1].ID + 1
	}
	d.docs = append(d.docs, doc)
	d.mu.Unlock()
}

func (d *DB) Delete(id int) error {
	i := sort.Search(len(d.docs), func(i int) bool { return d.docs[i].ID >= id })
	if d.docs[i].ID != id {
		return errors.New(fmt.Sprint("Not found doc ID:%v", id))
	}
	d.mu.Lock()
	switch len(d.docs) {
	case 1:
		d.docs = []crawler.Document{}
	default:
		d.docs = append(d.docs[0:i], d.docs[i+1:]...)
	}
	d.mu.Unlock()
	return nil
}

func (d *DB) Update(doc crawler.Document) error {
	i := sort.Search(len(d.docs), func(i int) bool { return d.docs[i].ID >= doc.ID })
	if d.docs[i].ID != doc.ID {
		return errors.New(fmt.Sprintf("Not found doc ID:%v", doc.ID))
	}
	d.mu.Lock()
	if doc.Title != "" {
		d.docs[i].Title = doc.Title
	}
	if doc.URL != "" {
		d.docs[i].URL = doc.URL
	}
	d.mu.Unlock()
	return nil
}
