package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jcr04/Zapfast.go/app/pkg/usecase"
)

type RideHandler struct {
	rideUsecase usecase.RideUsecase
}

func NewRideHandler(ru usecase.RideUsecase) *RideHandler {
	return &RideHandler{rideUsecase: ru}
}

func (rh *RideHandler) RequestRide(w http.ResponseWriter, r *http.Request) {
	// Aqui você pode parsear os dados da requisição, como customerID, startLocation, endLocation e price
	ride, err := rh.rideUsecase.RequestRide( /* argumentos */ )
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(ride)
}

func (rh *RideHandler) AcceptRide(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rideID, _ := strconv.Atoi(vars["rideID"])
	driverID, _ := strconv.Atoi(vars["driverID"])
	err := rh.rideUsecase.AcceptRide(rideID, driverID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (rh *RideHandler) CompleteRide(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rideID, _ := strconv.Atoi(vars["rideID"])
	err := rh.rideUsecase.CompleteRide(rideID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
