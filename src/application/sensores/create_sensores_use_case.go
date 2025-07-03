package sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type CreateSensoresUseCase struct {
	repo repositories.SensoresRepositories
}

func NewCreateSensoresUseCase(repo repositories.SensoresRepositories) *CreateSensoresUseCase {
	return &CreateSensoresUseCase{
		repo: repo,
	}
}

func (uc *CreateSensoresUseCase) Execute(sensor entities.Sensores) error {
	return uc.repo.CreateSensor(sensor)
}
