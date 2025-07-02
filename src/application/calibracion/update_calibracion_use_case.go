package calibracion

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type UpdateCalibracionUseCase struct {
	repo repositories.CalibracionRepositories
}

func NewUpdateCalibracionUseCase(repo repositories.CalibracionRepositories) *UpdateCalibracionUseCase {
	return &UpdateCalibracionUseCase{
		repo: repo,
	}
}

func (uc *UpdateCalibracionUseCase) Execute(calibracion entities.Calibracion) error {
	return uc.repo.UpdateCalibracion(calibracion)
}
