package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"thinknetica/13-api/pkg/crawler"
	"thinknetica/13-api/pkg/storage"
)

type API struct {
	storage *storage.Storage
	router  *mux.Router
}

func New(s *storage.Storage) *API {
	a := API{}
	a.router = mux.NewRouter()
	a.storage = s
	a.endpoints()
	return &a
}

func (a *API) endpoints() {
	a.router.HandleFunc("/docs", a.AllDocs).Methods(http.MethodGet, http.MethodOptions)
	a.router.HandleFunc("/docs/{keyword}", a.Search).Methods(http.MethodGet, http.MethodOptions)
	a.router.HandleFunc("/docs/{id}", a.Delete).Methods(http.MethodDelete, http.MethodOptions)
	a.router.HandleFunc("/docs", a.Add).Methods(http.MethodPost, http.MethodOptions)
	a.router.HandleFunc("/docs/{id}", a.Update).Methods(http.MethodPatch, http.MethodOptions)
}

func (a *API) Router() *mux.Router {
	return a.router
}

func (a *API) AllDocs(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(*a.storage.All())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *API) Search(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	keyword := p["keyword"]
	if keyword == "" {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	docs := a.storage.Search(keyword)
	err := json.NewEncoder(w).Encode(docs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id, err := strconv.Atoi(p["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = a.storage.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *API) Add(w http.ResponseWriter, r *http.Request) {
	var doc crawler.Document
	err := json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a.storage.DB.Add(doc)
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id, err := strconv.Atoi(p["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var doc crawler.Document
	err = json.NewDecoder(r.Body).Decode(&doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	doc.ID = id
	err = a.storage.Update(doc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
