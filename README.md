# fibo
A fibonacci fun house

This application exposes three endpoints

Prerequisite: Please run postgres https://www.postgresql.org/docs/
Step 1. run tests `go test ./... -v`
Step 2. run application `go run main.go`

3 endpoints are exposed
GET localhost:8000/fibonacci/{number}
GET localhost:8000/memoizedresults/{memoNumber}
GET localhost:8000/deleteall

Note: Delete is usually a delete request but I made it a Get for simplicity
