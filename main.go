package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	print("Starting Server at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to golang world!")
}
