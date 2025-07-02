package calibracion

import (
	"apisensores/src/domain/repositories"
)

type DeleteCalibracionUseCase struct {
	repo repositories.CalibracionRepositories
}

func NewDeleteCalibracionUseCase(repo repositories.CalibracionRepositories) *DeleteCalibracionUseCase {
	return &DeleteCalibracionUseCase{
		repo: repo,
	}
}

func (uc *DeleteCalibracionUseCase) Execute(id int) error {
	return uc.repo.DeleteCalibracion(id)
}
