package index

import (
	"strings"
	"thinknetica/10-batchscan/pkg/crawler"
)

type Index struct {
	data map[string][]int
}

func (r *Index) New() {
	r.data = make(map[string][]int)
}

func (i *Index) Add(title string, id int) {
	keys := strings.Split(strings.ToLower(title), " ")
	for _, key := range keys {
		if len(i.data[key]) == 0 || i.data[key][len(i.data[key])-1] != id {
			i.data[key] = append(i.data[key], id)
		}
	}
}

func (i *Index) Search(keyword string) []int {
	return i.data[strings.ToLower(keyword)]
}

func (i *Index) Create(docs []crawler.Document) {
	i.New()
	for _, elem := range docs {
		i.Add(elem.Title, elem.ID)
	}
}
