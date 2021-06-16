package handlers

import (
	"io"
	"net/http"
)

func FibHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to golang world!")
}

func MemoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to golang world!")
}
