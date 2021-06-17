package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/fibo/database"
	"github.com/gorilla/mux"
)

func DeleteAll(w http.ResponseWriter, r *http.Request) {
	e := database.DeleteAll()
	if e != nil {
		io.WriteString(w, fmt.Sprint(e.Error()))
	}
	io.WriteString(w, "success, all deleted")
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	number := params["number"]
	i, _ := strconv.Atoi(number)
	fibNumber := GetFibonacci(i)
	io.WriteString(w, fmt.Sprint(fibNumber))
}

func MemoHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	m := params["memoNumber"]
	i, _ := strconv.Atoi(m)
	var memoNumber int
	memoNumber, err := database.GetMemo(i)
	if err != nil {
		memoNumber = GetMemoizationNumber(i)
		database.InsertRow(i, memoNumber)
	}
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
