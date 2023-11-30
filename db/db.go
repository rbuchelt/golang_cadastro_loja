package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectBd() *sql.DB {
	conn := "user=dev dbname=loja password=ituran host=127.0.0.1 sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
