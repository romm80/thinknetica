package webapp

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"thinknetica/12-http/pkg/crawler"
)

var wa *webapp

func TestMain(m *testing.M) {
	wa = new(webapp)
	wa.docs = []crawler.Document{}
	wa.index = make(map[string][]int)
	wa.Router = mux.NewRouter()
	endpoints(wa)
	os.Exit(m.Run())
}
func Test_webapp_mainHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	wa.Router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
	exepted := `<html><body><p><a href="./index">index</a></p><p><a href="./docs">docs</a></p></body></html>`
	if rr.Body.String() != exepted {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), exepted)
	}
}

func Test_webapp_getDocs(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/docs", nil)
	rr := httptest.NewRecorder()

	wa.Router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
	exepted, _ := json.Marshal(wa.docs)
	if rr.Body.String() != string(exepted) {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), string(exepted))
	}
}

func Test_webapp_getIndex(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/index", nil)
	rr := httptest.NewRecorder()

	wa.Router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}
	exepted, _ := json.Marshal(wa.index)
	if rr.Body.String() != string(exepted) {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), string(exepted))
	}
}
