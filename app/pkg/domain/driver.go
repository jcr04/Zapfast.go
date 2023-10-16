package domain

import "time"

type Driver struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
