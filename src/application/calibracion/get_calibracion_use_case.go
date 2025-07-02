package calibracion

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetCalibracionUseCase struct {
	repo repositories.CalibracionRepositories
}

func NewGetCalibracionUseCase(repo repositories.CalibracionRepositories) *GetCalibracionUseCase {
	return &GetCalibracionUseCase{
		repo: repo,
	}
}

func (uc *GetCalibracionUseCase) Execute(id int) (entities.Calibracion, error) {
	return uc.repo.GetCalibracion(id)
}
