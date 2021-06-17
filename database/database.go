package database

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "hide"
	dbname   = "memo"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func InitDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	checkError(err)

	defer db.Close()

	err = db.Ping()
	checkError(err)

	fmt.Println("Connected!")
}
