// C:\Users\Usuario\go\Zapfast.go\app\pkg\repository\api\api.go

package api

import (
	"database/sql"

	"github.com/gorilla/mux"
	customHttp "github.com/jcr04/Zapfast.go/app/pkg/delivery/http"
	"github.com/jcr04/Zapfast.go/app/pkg/repository"
	"github.com/jcr04/Zapfast.go/app/pkg/usecase"
)

func NewRouter(db *sql.DB) *mux.Router {
	// Repositórios
	rideRepo := repository.NewRideRepository(db)
	customerRepo := repository.NewPostgresCustomerRepository(db)
	driverRepo := repository.NewDriverRepository(db)

	// Casos de uso
	rideUsecase := usecase.NewRideUsecase(rideRepo)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	driverUsecase := usecase.NewDriverUsecase(driverRepo)

	// Handlers
	rideHandler := customHttp.NewRideHandler(rideUsecase)
	customerHandler := customHttp.NewCustomerHandler(customerUsecase)
	driverHandler := customHttp.NewDriverHandler(driverUsecase)

	// Configuração do router
	r := mux.NewRouter()
	r.HandleFunc("/rides", rideHandler.RequestRide).Methods("POST")
	r.HandleFunc("/rides/{rideID}/accept", rideHandler.AcceptRide).Methods("PATCH")
	r.HandleFunc("/rides/{rideID}/complete", rideHandler.CompleteRide).Methods("PATCH")

	// Endpoints de Customer
	r.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	r.HandleFunc("/customers/{customerID}", customerHandler.GetCustomer).Methods("GET")
	r.HandleFunc("/customers/{customerID}", customerHandler.UpdateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{customerID}", customerHandler.DeleteCustomer).Methods("DELETE")

	// Endpoints de Driver
	r.HandleFunc("/drivers", driverHandler.CreateDriver).Methods("POST")
	r.HandleFunc("/drivers/{driverID}", driverHandler.GetDriver).Methods("GET")
	r.HandleFunc("/drivers/{driverID}", driverHandler.UpdateDriver).Methods("PUT")
	r.HandleFunc("/drivers/{driverID}", driverHandler.DeleteDriver).Methods("DELETE")

	return r
}
