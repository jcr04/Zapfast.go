package usecase

import (
	"github.com/jcr04/Zapfast.go/app/pkg/domain"
	"github.com/jcr04/Zapfast.go/app/pkg/repository"
)

type CustomerUsecase struct {
	repo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) *CustomerUsecase {
	return &CustomerUsecase{repo: repo}
}

func (cu *CustomerUsecase) CreateCustomer(customer domain.Customer) (domain.Customer, error) {
	return cu.repo.CreateCustomer(customer)
}

func (cu *CustomerUsecase) GetCustomer(id int) (domain.Customer, error) {
	return cu.repo.GetCustomer(id)
}

func (cu *CustomerUsecase) UpdateCustomer(customer domain.Customer) (domain.Customer, error) {
	return cu.repo.UpdateCustomer(customer)
}

func (cu *CustomerUsecase) DeleteCustomer(id int) error {
	return cu.repo.DeleteCustomer(id)
}
