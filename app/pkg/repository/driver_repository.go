// C:\Users\Usuario\go\Zapfast.go\app\pkg\repository\driver_repository.go

package repository

import (
	"database/sql"

	"github.com/jcr04/Zapfast.go/app/pkg/domain"
)

type DriverRepository struct {
	db *sql.DB
}

func NewDriverRepository(db *sql.DB) *DriverRepository {
	return &DriverRepository{db: db}
}

func (dr *DriverRepository) CreateDriver(driver domain.Driver) (domain.Driver, error) {
	query := `INSERT INTO drivers (name, email, phone) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at;`
	err := dr.db.QueryRow(query, driver.Name, driver.Email, driver.Phone).Scan(&driver.ID, &driver.CreatedAt, &driver.UpdatedAt)
	if err != nil {
		return domain.Driver{}, err
	}
	return driver, nil
}

func (dr *DriverRepository) GetDriver(id int) (domain.Driver, error) {
	var driver domain.Driver
	query := `SELECT * FROM drivers WHERE id = $1;`
	err := dr.db.QueryRow(query, id).Scan(&driver.ID, &driver.Name, &driver.Email, &driver.Phone, &driver.CreatedAt, &driver.UpdatedAt)
	if err != nil {
		return domain.Driver{}, err
	}
	return driver, nil
}

func (dr *DriverRepository) UpdateDriver(driver domain.Driver) (domain.Driver, error) {
	query := `UPDATE drivers SET name = $1, email = $2, phone = $3, updated_at = NOW() WHERE id = $4 RETURNING updated_at;`
	err := dr.db.QueryRow(query, driver.Name, driver.Email, driver.Phone, driver.ID).Scan(&driver.UpdatedAt)
	if err != nil {
		return domain.Driver{}, err
	}
	return driver, nil
}

func (dr *DriverRepository) DeleteDriver(id int) error {
	query := `DELETE FROM drivers WHERE id = $1;`
	_, err := dr.db.Exec(query, id)
	return err
}
