package sensores

import (
	"apisensores/src/domain/repositories"
)

type DeleteSensoresUseCase struct {
	repo repositories.SensoresRepositories
}

func NewDeleteSensoresUseCase(repo repositories.SensoresRepositories) *DeleteSensoresUseCase {
	return &DeleteSensoresUseCase{
		repo: repo,
	}
}

func (uc *DeleteSensoresUseCase) Execute(id int) error {
	return uc.repo.DeleteSensor(id)
}
