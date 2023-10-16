package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	connStr := "user=postgres password=12345 dbname=zapfast host=localhost port=5432 sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// GetDB retorna uma conexão com o banco de dados
func GetDB() *sql.DB {
	return db
}
