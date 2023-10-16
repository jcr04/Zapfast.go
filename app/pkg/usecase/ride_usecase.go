package usecase

import (
	"github.com/jcr04/Zapfast.go/app/pkg/domain"
	"github.com/jcr04/Zapfast.go/app/pkg/repository"
)

type RideUsecase struct {
	rideRepo repository.RideRepository
}

func NewRideUsecase(rr repository.RideRepository) *RideUsecase {
	return &RideUsecase{rideRepo: rr}
}

func (ru *RideUsecase) RequestRide(customerID int, startLocation, endLocation string, price float64) (*domain.Ride, error) {
	ride := &domain.Ride{
		CustomerID:    customerID,
		StartLocation: startLocation,
		EndLocation:   endLocation,
		Status:        domain.StatusRequested,
		Price:         price,
	}
	return ru.rideRepo.Store(ride)
}

func (ru *RideUsecase) AcceptRide(rideID, driverID int) error {
	ride, err := ru.rideRepo.FindByID(rideID)
	if err != nil {
		return err
	}
	ride.Accept(driverID)
	return ru.rideRepo.Update(ride)
}

func (ru *RideUsecase) CompleteRide(rideID int) error {
	ride, err := ru.rideRepo.FindByID(rideID)
	if err != nil {
		return err
	}
	ride.Complete()
	return ru.rideRepo.Update(ride)
}
