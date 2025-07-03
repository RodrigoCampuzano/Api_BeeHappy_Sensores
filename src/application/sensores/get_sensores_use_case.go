package sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetSensoresUseCase struct {
	repo repositories.SensoresRepositories	
}

func NewGetSensoresUseCase(repo repositories.SensoresRepositories) *GetSensoresUseCase {
	return &GetSensoresUseCase{
		repo: repo,
	}
}

func (uc *GetSensoresUseCase) Execute(id int) (entities.Sensores, error) {
	return uc.repo.GetSensor(id)
}
