package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fibonacci", FibHandler)
	r.HandleFunc("/memoizedresults", MemoHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to golang world!")
}

func MemoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to golang world!")
}
