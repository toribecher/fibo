package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/fibo/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hide"
	dbname   = "memo"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")
	r := mux.NewRouter()
	r.HandleFunc("/fibonacci/{number}", handlers.FibHandler)
	r.HandleFunc("/memoizedresults", handlers.MemoHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
