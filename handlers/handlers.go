package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func FibHandler(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	number := split[len(split)-1]
	i, _ := strconv.Atoi(number)

	fibNumber := FibonacciRecursion(i)
	io.WriteString(w, fmt.Sprint(fibNumber))
}

func MemoHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "welcome to golang world!")
}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
