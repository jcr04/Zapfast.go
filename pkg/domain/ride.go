package domain

import "time"

type RideStatus string

const (
	StatusRequested RideStatus = "REQUESTED"
	StatusAccepted  RideStatus = "ACCEPTED"
	StatusCompleted RideStatus = "COMPLETED"
)

type Ride struct {
	ID            int
	CustomerID    int
	DriverID      int
	StartLocation string
	EndLocation   string
	Status        RideStatus
	Price         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Método para aceitar uma corrida
func (r *Ride) Accept(driverID int) {
	r.DriverID = driverID
	r.Status = StatusAccepted
}

// Método para completar uma corrida
func (r *Ride) Complete() {
	r.Status = StatusCompleted
}
