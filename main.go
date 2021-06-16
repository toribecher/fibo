package main

import (
	"log"
	"net/http"

	"github.com/fibo/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fibonacci", handlers.FibHandler)
	r.HandleFunc("/memoizedresults", handlers.MemoHandler)
	// http.go Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
