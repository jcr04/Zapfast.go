// C:\Users\Usuario\go\Zapfast.go\app\cmd\server\main.go

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jcr04/Zapfast.go/app/pkg/repository/api"
)

func main() {
	connStr := "user=postgres password=12345 dbname=zapfast host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	r := api.NewRouter(db)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
