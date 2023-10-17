// C:\Users\Usuario\go\Zapfast.go\app\pkg\usecase\driver_usecase.go

package usecase

import (
	"github.com/jcr04/Zapfast.go/app/pkg/domain"
	"github.com/jcr04/Zapfast.go/app/pkg/repository"
)

type DriverUsecase struct {
	repo repository.DriverRepository
}

func NewDriverUsecase(repo repository.DriverRepository) *DriverUsecase {
	return &DriverUsecase{repo: repo}
}

func (du *DriverUsecase) CreateDriver(driver domain.Driver) (domain.Driver, error) {
	return du.repo.CreateDriver(driver)
}

func (du *DriverUsecase) GetDriver(id int) (domain.Driver, error) {
	return du.repo.GetDriver(id)
}

func (du *DriverUsecase) UpdateDriver(driver domain.Driver) (domain.Driver, error) {
	return du.repo.UpdateDriver(driver)
}

func (du *DriverUsecase) DeleteDriver(id int) error {
	return du.repo.DeleteDriver(id)
}
