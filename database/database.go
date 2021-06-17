package database

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	db       *sql.DB
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = os.Getenv("PASSWORD")
	dbname   = "memo"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlconn)
	checkError(err)

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")
}

func InsertRow(number, memoCount int) {
	insertDynStmt := `INSERT into "memos"("number", "count") values($1, $2)`
	_, e := db.Exec(insertDynStmt, number, memoCount)
	checkError(e)
}

func GetMemo(number int) (int, error) {
	var memoCount int
	getDynStmt := `SELECT count FROM "memos" WHERE number=$1`
	row := db.QueryRow(getDynStmt, number)
	err := row.Scan(&memoCount)
	return memoCount, err
}

func DeleteAll() error {
	deleteStmt := `DELETE FROM "memos"`
	_, e := db.Exec(deleteStmt)
	return e
}
