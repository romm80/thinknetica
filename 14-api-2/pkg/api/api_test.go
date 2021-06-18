package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI_authentication(t *testing.T) {
	api := New()
	log := LogInfo{
		Usr: "usr1",
		Psw: "qwerty",
	}
	js, _ := json.Marshal(log)
	req := httptest.NewRequest(http.MethodPost, "/auth", bytes.NewBuffer(js))
	rr := httptest.NewRecorder()
	api.router.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Error("Got #{rr.Code} - want #{http.StatusOK}")
	}
	t.Log(rr.Body)
}
