package repository

import (
	"database/sql"

	"github.com/jcr04/Zapfast.go/app/pkg/domain"
)

type CustomerRepository interface {
	CreateCustomer(domain.Customer) (domain.Customer, error)
	GetCustomer(int) (domain.Customer, error)
	UpdateCustomer(domain.Customer) (domain.Customer, error)
	DeleteCustomer(int) error
}

type PostgresCustomerRepository struct {
	db *sql.DB
}

func NewPostgresCustomerRepository(db *sql.DB) *PostgresCustomerRepository {
	return &PostgresCustomerRepository{db: db}
}

func (repo *PostgresCustomerRepository) CreateCustomer(customer domain.Customer) (domain.Customer, error) {
	query := "INSERT INTO customers (name, email, phone) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at"
	err := repo.db.QueryRow(query, customer.Name, customer.Email, customer.Phone).Scan(&customer.ID, &customer.CreatedAt, &customer.UpdatedAt)
	return customer, err
}

func (repo *PostgresCustomerRepository) GetCustomer(id int) (domain.Customer, error) {
	var customer domain.Customer
	query := "SELECT id, name, email, phone, created_at, updated_at FROM customers WHERE id = $1"
	err := repo.db.QueryRow(query, id).Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Phone, &customer.CreatedAt, &customer.UpdatedAt)
	return customer, err
}

func (repo *PostgresCustomerRepository) UpdateCustomer(customer domain.Customer) (domain.Customer, error) {
	query := "UPDATE customers SET name = $1, email = $2, phone = $3, updated_at = NOW() WHERE id = $4 RETURNING updated_at"
	err := repo.db.QueryRow(query, customer.Name, customer.Email, customer.Phone, customer.ID).Scan(&customer.UpdatedAt)
	return customer, err
}

func (repo *PostgresCustomerRepository) DeleteCustomer(id int) error {
	query := "DELETE FROM customers WHERE id = $1"
	_, err := repo.db.Exec(query, id)
	return err
}
