package ws

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type API struct {
	router   *mux.Router
	upgrader websocket.Upgrader
	chIn     chan string
	chOut    []chan string
	mu       sync.Mutex
}

func New() *API {
	a := API{}
	a.router = mux.NewRouter()
	a.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	a.chIn = make(chan string)
	a.endpoints()
	return &a
}

func (a *API) Router() *mux.Router {
	return a.router
}

func (a *API) endpoints() {
	a.router.HandleFunc("/messages", a.messages)
	a.router.HandleFunc("/send", a.send)
}

func (a *API) messages(w http.ResponseWriter, r *http.Request) {
	conn, err := a.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	a.mu.Lock()
	out := make(chan string)
	a.chOut = append(a.chOut, out)
	a.mu.Unlock()
	defer func() {
		a.mu.Lock()
		for i, ch := range a.chOut {
			if ch == out {
				a.chOut = append(a.chOut[:i], a.chOut[i+1:]...)
			}
		}
		a.mu.Unlock()
	}()

	for msg := range out {
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println(err.Error())
			return
		}
	}
}

func (a *API) send(w http.ResponseWriter, r *http.Request) {
	conn, err := a.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}
		a.chIn <- string(message)
	}
}

func (a *API) ToClients() {
	for in := range a.chIn {
		for _, out := range a.chOut {
			out <- in
		}
	}
}
