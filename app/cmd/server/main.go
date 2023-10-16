package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	customHttp "github.com/jcr04/Zapfast.go/app/pkg/delivery/http"
	repository "github.com/jcr04/Zapfast.go/app/pkg/repository/postgres"
	"github.com/jcr04/Zapfast.go/app/pkg/usecase"
)

func main() {
	connStr := "user=postgres password=12345 dbname=zapfast host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rideRepo := repository.NewRideRepository(db)
	rideUsecase := usecase.NewRideUsecase(*rideRepo)
	rideHandler := customHttp.NewRideHandler(rideUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/rides", rideHandler.RequestRide).Methods("POST")
	r.HandleFunc("/rides/{rideID}/accept", rideHandler.AcceptRide).Methods("PATCH")
	r.HandleFunc("/rides/{rideID}/complete", rideHandler.CompleteRide).Methods("PATCH")

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
