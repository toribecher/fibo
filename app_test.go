package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		host,
		port,
		user,
		password,
		dbname)

	ensureTableExists()
	code := m.Run()
	// clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}
func TestGetMemoTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/memoizedresults/120", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"number":120,"memoNumber":12}` {
		t.Errorf("Expected 120 and 12 Got %s", body)
	}
}

func TestEmptyMemoTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/memoizedresults/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestZeroMemoTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/memoizedresults/0", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"number":0,"memoNumber":0}` {
		t.Errorf("Expected 0. Got %s", body)
	}
}

func TestOneMemoTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/memoizedresults/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"number":1,"memoNumber":1}` {
		t.Errorf("Expected 1. Got %s", body)
	}
}

func TestTwoMemoTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/memoizedresults/2", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"number":2,"memoNumber":3}` {
		t.Errorf("Expected 1. Got %s", body)
	}
}

func TestMemoAlreadyExists(t *testing.T) {
	clearTable()

	m := memo{Number: 3, MemoNumber: 4}
	m.createMemo(a.DB)

	req, _ := http.NewRequest("GET", "/memoizedresults/3", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != `{"number":3,"memoNumber":4}` {
		t.Errorf("Expected 3, 4. Got %s", body)
	}
}
func clearTable() {
	a.DB.Exec("DELETE FROM memos")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetFibo(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/11", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `89` {
		t.Errorf("Expected 89, Got %s", body)
	}
}

func TestGetFibo2(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/12", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `144` {
		t.Errorf("Expected 89, Got %s", body)
	}
}

func TestEmptyFibo(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetFibo0(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/0", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `0` {
		t.Errorf("Expected 0, Got %s", body)
	}
}

func TestGetFibo1(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `1` {
		t.Errorf("Expected 1, Got %s", body)
	}
}

func TestGetFibo02(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/2", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `1` {
		t.Errorf("Expected 1, Got %s", body)
	}
}

func TestGetFibo03(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/fibonacci/3", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	if body := response.Body.String(); body != `2` {
		t.Errorf("Expected 2, Got %s", body)
	}
}

func TestDeleteMemos(t *testing.T) {
	clearTable()
	m := memo{Number: 3, MemoNumber: 4}
	m.createMemo(a.DB)

	m2 := memo{Number: 1, MemoNumber: 1}
	m2.createMemo(a.DB)

	m3 := memo{Number: 0, MemoNumber: 0}
	m3.createMemo(a.DB)

	m4 := memo{Number: 11, MemoNumber: 89}
	m4.createMemo(a.DB)

	req, _ := http.NewRequest("GET", "/memoizedresults/11", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/deleteall", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	rows, err := getAllMemos(a.DB)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	if count != 0 {
		t.Errorf("Expected 0, Got %d", count)
	}

}

func getAllMemos(db *sql.DB) (*sql.Rows, error) {
	return db.Query(`SELECT COUNT(*) FROM "memos"`)
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS memos
(
   number INT PRIMARY KEY,
   memoNumber INT NOT NULL
)`
