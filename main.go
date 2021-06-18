package main

import (
	"os"

	"github.com/fibo/app"
)

var (
	host     = os.Getenv("HOST")
	port     = os.Getenv("PORT")
	user     = os.Getenv("USER")
	password = os.Getenv("PASSWORD")
	dbname   = os.Getenv("DBNAME")
)

func main() {
	a := app.App{}

	a.Initialize(
		host,
		port,
		user,
		password,
		dbname)

	a.Run(":8000")
}
