package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectBD() *sql.DB {
	con := "user=postgres dbname=LojaGo password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		panic(err.Error())
	} else {
		return db
	}
}
