package main

import (
	"database/sql"
)

type memo struct {
	Number     int `json:"number"`
	MemoNumber int `json:"memoNumber"`
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
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
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

func (m *memo) getMemo(db *sql.DB) error {
	return db.QueryRow(`SELECT count FROM "memos" WHERE number=$1`, m.Number).Scan(&m.MemoNumber)
}

func (m *memo) createMemo(db *sql.DB) error {
	insertDynStmt := `INSERT into "memos"("number", "count") values($1, $2)`
	_, e := db.Exec(insertDynStmt, m.Number, m.MemoNumber)
	return e
}

func DeleteAll(db *sql.DB) error {
	deleteStmt := `DELETE FROM "memos"`
	_, e := db.Exec(deleteStmt)
	return e
}
