package index

import (
	"sort"
	"strings"
	"thinknetica/05-io/pkg/crawler"
)

type Index struct {
	data map[string][]int
}

func (r *Index) New() {
	r.data = make(map[string][]int)
}

func (r *Index) Add(title string, i int) {
	keys := strings.Split(strings.ToLower(title), " ")
	for _, key := range keys {
		if len(r.data[key]) == 0 || r.data[key][len(r.data[key]) - 1] != i {
			r.data[key] = append(r.data[key], i)
		}
	}
}

func (r *Index) Search(keyword string) []int {
	return r.data[strings.ToLower(keyword)]
}

func (r *Index) Create(docs []crawler.Document) {
	r.New()
	for _, elem := range docs {
		r.Add(elem.Title, elem.ID)
	}
	sort.Slice(docs, func(i, j int) bool { return docs[i].ID <= docs[j].ID })
}