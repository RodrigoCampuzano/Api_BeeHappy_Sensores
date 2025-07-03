package sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type UpdateSensoresUseCase struct {
	repo repositories.SensoresRepositories
}

func NewUpdateSensoresUseCase(repo repositories.SensoresRepositories) *UpdateSensoresUseCase {
	return &UpdateSensoresUseCase{
		repo: repo,
	}
}

func (uc *UpdateSensoresUseCase) Execute(sensor entities.Sensores) error {
	return uc.repo.UpdateSensor(sensor)
}
