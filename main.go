package main

import (
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = os.Getenv("PASSWORD")
	dbname   = os.Getenv("DBNAME")
)

func main() {
	a := App{}
	a.Initialize(
		host,
		port,
		user,
		password,
		dbname)

	a.Run(":8010")
}

// func main() {
// 	database.InitDB()
// 	r := mux.NewRouter()
// 	r.HandleFunc("/delete", handlers.DeleteAll)
// 	r.HandleFunc("/fibonacci/{number}", handlers.FibHandler)
// 	r.HandleFunc("/memoizedresults/{memoNumber}", handlers.MemoHandler)
// 	log.Fatal(http.ListenAndServe(":8000", r))
// }
