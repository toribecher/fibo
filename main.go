package main

import (
	"log"
	"net/http"

	"github.com/fibo/database"
	"github.com/fibo/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	database.InitDB()
	r := mux.NewRouter()
	r.HandleFunc("/fibonacci/{number}", handlers.FibHandler)
	r.HandleFunc("/memoizedresults/{memoNumber}", handlers.MemoHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
