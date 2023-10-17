package repository

import (
	"database/sql"

	"github.com/jcr04/Zapfast.go/app/pkg/domain"
)

type RideRepository struct {
	DB *sql.DB
}

func NewRideRepository(db *sql.DB) *RideRepository {
	return &RideRepository{DB: db}
}

func (rr *RideRepository) Store(ride *domain.Ride) (*domain.Ride, error) {
	query := `
		INSERT INTO rides (customer_id, driver_id, start_location, end_location, status, price)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at;
	`
	err := rr.DB.QueryRow(query, ride.CustomerID, ride.DriverID, ride.StartLocation, ride.EndLocation, ride.Status, ride.Price).Scan(&ride.ID, &ride.CreatedAt, &ride.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return ride, nil
}

func (rr *RideRepository) FindByID(id int) (*domain.Ride, error) {
	ride := &domain.Ride{}
	query := `
		SELECT * FROM rides WHERE id = $1;
	`
	err := rr.DB.QueryRow(query, id).Scan(&ride.ID, &ride.CustomerID, &ride.DriverID, &ride.StartLocation, &ride.EndLocation, &ride.Status, &ride.Price, &ride.CreatedAt, &ride.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return ride, nil
}

func (rr *RideRepository) Update(ride *domain.Ride) error {
	query := `
		UPDATE rides
		SET driver_id = $2, status = $3, updated_at = NOW()
		WHERE id = $1;
	`
	_, err := rr.DB.Exec(query, ride.ID, ride.DriverID, ride.Status)
	return err
}
