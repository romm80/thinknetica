package index

import (
	"strings"
	"thinknetica/13-api/pkg/crawler"
)

type Index struct {
	data map[string][]int
}

func New(docs []crawler.Document) *Index {
	i := Index{make(map[string][]int)}
	for _, doc := range docs {
		i.Add(doc)
	}
	return &i
}

func (i *Index) Add(doc crawler.Document) {
	keys := strings.Split(strings.ToLower(doc.Title), " ")
	for _, key := range keys {
		if len(i.data[key]) == 0 || i.data[key][len(i.data[key])-1] != doc.ID {
			i.data[key] = append(i.data[key], doc.ID)
		}
	}
}

func (i *Index) Search(keyword string) []int {
	return i.data[strings.ToLower(keyword)]
}
