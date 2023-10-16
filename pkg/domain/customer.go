package domain

import "time"

type Customer struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
