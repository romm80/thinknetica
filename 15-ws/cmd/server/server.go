package main

import (
	"log"
	"net/http"
	"thinknetica/15-ws/pkg/ws"
)

func main() {
	api := ws.New()
	go api.ToClients()
	log.Fatal(http.ListenAndServe(":8080", api.Router()))
}
