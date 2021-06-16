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

	fibNumber := GetFibonacci(i)
	io.WriteString(w, fmt.Sprint(fibNumber))
}

func MemoHandler(w http.ResponseWriter, r *http.Request) {
	split := strings.Split(r.URL.Path, "/")
	number := split[len(split)-1]
	i, _ := strconv.Atoi(number)

	memoNumber := GetMemoizationNumber(i)
	io.WriteString(w, fmt.Sprint(memoNumber))
}

func GetFibonacci(n int) int {
	fibBox := []int{0, 1}
	for i := 0; i < n; i++ {
		v := fibBox[i] + fibBox[i+1]
		fibBox = append(fibBox, v)
	}
	result := int(fibBox[n])

	return result
}

func GetMemoizationNumber(n int) int {
	fibBox := []int{0, 1}
	for i := 0; i < n; i++ {
		v := fibBox[i] + fibBox[i+1]
		fibBox = append(fibBox, v)
		if v >= n {
			break
		}
	}
	result := int(len(fibBox) - 1)

	return result
}
