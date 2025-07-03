package sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetAllSensoresUseCase struct {
	repo repositories.SensoresRepositories
}

func NewGetAllSensoresUseCase(repo repositories.SensoresRepositories) *GetAllSensoresUseCase {
	return &GetAllSensoresUseCase{
		repo: repo,
	}
}

func (uc *GetAllSensoresUseCase) Execute() ([]entities.Sensores, error) {
	return uc.repo.GetAllSensores()
}
