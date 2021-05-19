package index

import "strings"

type Index struct {
	data map[string][]int
}

func (r *Index) New() {
	r.data = make(map[string][]int)
}

func (r *Index) Add(title string, i int) {
	keys := strings.Split(strings.ToLower(title), " ")
	for _, key := range keys {
		r.data[key] = append(r.data[key], i)
	}
}

func (r *Index) Search(keyword string) []int {
	return r.data[strings.ToLower(keyword)]
}