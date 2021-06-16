package storage

import (
	"thinknetica/13-api/pkg/crawler"
	"thinknetica/13-api/pkg/storage/db"
	"thinknetica/13-api/pkg/storage/index"
)

type Index interface {
	Search(keyword string) []int
}

type DB interface {
	All() *[]crawler.Document
	Search(ids []int) []crawler.Document
	Add(doc crawler.Document)
	Delete(id int) error
	Update(doc crawler.Document) error
}

type Storage struct {
	DB
	Index
}

func (s *Storage) Search(keyword string) []crawler.Document {
	return s.DB.Search(s.Index.Search(keyword))
}

func New(docs []crawler.Document) *Storage {
	return &Storage{
		DB:    db.New(docs),
		Index: index.New(docs),
	}
}
