package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(host string, port int, user, password, dbname string) {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/deleteall", a.deleteAll).Methods("DELETE")
	// a.Router.HandleFunc("/products", a.getFibonacci).Methods("GET")

	a.Router.HandleFunc("/fibonacci/{number}", a.getFibonacci).Methods("GET")
	a.Router.HandleFunc("/memoizedresults/{memoNumber}", a.memoHandler).Methods("GET")
}

func (a *App) deleteAll(w http.ResponseWriter, r *http.Request) {
	err := DeleteAll(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	fmt.Println("SUCCESS")
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) getFibonacci(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	number := params["number"]
	i, _ := strconv.Atoi(number)
	fibNumber := GetFibonacci(i)
	io.WriteString(w, fmt.Sprint(fibNumber))
}

func (a *App) memoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stringID := params["memoNumber"]
	id, _ := strconv.Atoi(stringID)
	memo := memo{Number: id}
	err := memo.getMemo(a.DB)
	if err != nil {
		memo.MemoNumber = GetMemoizationNumber(id)
		err = memo.createMemo(a.DB)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		}
	}
	respondWithJSON(w, http.StatusOK, memo)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
