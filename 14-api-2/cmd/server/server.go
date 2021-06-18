package main

import (
	"log"
	"net/http"
	"thinknetica/14-api-2/pkg/api"
)

func main() {
	a := api.New()
	log.Fatal(http.ListenAndServe(":8000", a.Router()))
}
