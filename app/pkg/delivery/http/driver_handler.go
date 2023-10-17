// C:\Users\Usuario\go\Zapfast.go\app\pkg\delivery\http\driver_handler.go
package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jcr04/Zapfast.go/app/pkg/domain"
	"github.com/jcr04/Zapfast.go/app/pkg/usecase"
)

type DriverHandler struct {
	driverUsecase usecase.DriverUsecase
}

func NewDriverHandler(du usecase.DriverUsecase) *DriverHandler {
	return &DriverHandler{driverUsecase: du}
}

// Create Driver
func (dh *DriverHandler) CreateDriver(w http.ResponseWriter, r *http.Request) {
	var driver domain.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver, err := dh.driverUsecase.CreateDriver(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(driver)
}

func (dh *DriverHandler) GetDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["driverID"]
	if !ok {
		http.Error(w, "Driver ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver, err := dh.driverUsecase.GetDriver(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(driver)
}

func (dh *DriverHandler) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	var driver domain.Driver
	if err := json.NewDecoder(r.Body).Decode(&driver); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	driver, err := dh.driverUsecase.UpdateDriver(driver)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(driver)
}

func (dh *DriverHandler) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["driverID"]
	if !ok {
		http.Error(w, "Driver ID is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dh.driverUsecase.DeleteDriver(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
