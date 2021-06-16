package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"thinknetica/13-api/pkg/crawler"
	"thinknetica/13-api/pkg/storage"
)

var api *API

func TestMain(m *testing.M) {
	docs := []crawler.Document{}
	docs = append(docs, crawler.Document{ID: 1, URL: "test", Title: "test"}, crawler.Document{ID: 2, URL: "test2", Title: "test2"})
	api = New(storage.New(docs))
	os.Exit(m.Run())
}

func Equal(a, b crawler.Document) bool {
	return a.ID == b.ID && a.URL == b.URL && a.Title == b.Title
}

func TestAPI_Add(t *testing.T) {
	want := crawler.Document{URL: "test3", Title: "test3"}
	payload, _ := json.Marshal(want)
	req := httptest.NewRequest(http.MethodPost, "/docs", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %v, а хотели %v", rr.Code, http.StatusOK)
	}
	docs := api.storage.DB.Search([]int{3})
	if len(docs) == 0 {
		t.Fatal("Документ не добавлен")
	}
	got := docs[0]
	want.ID = 3
	if !Equal(want, got) {
		t.Fatal("Документ не добавлен")
	}
}

func TestAPI_Delete(t *testing.T) {
	want := len(*api.storage.All()) - 1
	req := httptest.NewRequest(http.MethodDelete, "/docs/1", nil)
	rr := httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %v, а хотели %v", rr.Code, http.StatusOK)
	}
	got := len(*api.storage.All())
	if got != want {
		t.Fatal("Документ не удален")
	}
}

func TestAPI_Update(t *testing.T) {
	want := crawler.Document{URL: "test111", Title: "test111"}
	payload, _ := json.Marshal(want)
	req := httptest.NewRequest(http.MethodPatch, "/docs/2", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %v, а хотели %v", rr.Code, http.StatusOK)
	}
	got := api.storage.DB.Search([]int{2})[0]
	want.ID = 2
	if !Equal(want, got) {
		t.Fatal("Документ не добавлен")
	}
}

func TestAPI_Search(t *testing.T) {
	want := crawler.Document{ID: 2, URL: "test2", Title: "test2"}
	req := httptest.NewRequest(http.MethodGet, "/docs/test2", nil)
	rr := httptest.NewRecorder()
	api.Router().ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %v, а хотели %v", rr.Code, http.StatusOK)
	}
	got := []crawler.Document{}
	err := json.NewDecoder(rr.Body).Decode(&got)
	if err != nil || !Equal(want, got[0]) {
		t.Fatal("Документ не добавлен")
	}
}
