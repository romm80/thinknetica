package webapp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"thinknetica/12-http/pkg/crawler"
	"thinknetica/12-http/pkg/index"
)

type webapp struct {
	Router *mux.Router
	index  map[string][]int
	docs   []crawler.Document
}

func Init(index index.Index, docs []crawler.Document) *webapp {
	wa := new(webapp)
	wa.index = index.Data()
	wa.docs = docs
	wa.Router = mux.NewRouter()
	endpoints(wa)
	return wa
}

func endpoints(wa *webapp) {
	wa.Router.HandleFunc("/", wa.mainHandler).Methods(http.MethodGet)
	wa.Router.HandleFunc("/docs", wa.getDocs).Methods(http.MethodGet)
	wa.Router.HandleFunc("/index", wa.getIndex).Methods(http.MethodGet)
}

func (wa *webapp) mainHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(`<html><body><p><a href="./index">index</a></p><p><a href="./docs">docs</a></p></body></html>`))
}

func (wa *webapp) getDocs(w http.ResponseWriter, req *http.Request) {
	d, err := json.Marshal(wa.docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(d)
}

func (wa webapp) getIndex(w http.ResponseWriter, req *http.Request) {
	d, err := json.Marshal(wa.index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(d)
}
