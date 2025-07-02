package calibracion

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type CreateCalibracionUseCase struct {
	repo repositories.CalibracionRepositories
}

func NewCreateCalibracionUseCase(repo repositories.CalibracionRepositories) *CreateCalibracionUseCase {
	return &CreateCalibracionUseCase{
		repo: repo,
	}
}

func (uc *CreateCalibracionUseCase) Execute(calibracion entities.Calibracion) error {
	return uc.repo.CreateCalibracion(calibracion)
}
