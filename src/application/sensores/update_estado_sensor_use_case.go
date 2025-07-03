package sensores

import (
	"apisensores/src/domain/repositories"
)

type UpdateEstadoSensorUseCase struct {
	repo repositories.SensoresRepositories
}

func NewUpdateEstadoSensorUseCase(repo repositories.SensoresRepositories) *UpdateEstadoSensorUseCase {
	return &UpdateEstadoSensorUseCase{
		repo: repo,
	}
}

func (uc *UpdateEstadoSensorUseCase) Execute(id int, estado string) error {
	return uc.repo.UpdateEstadoSensor(id, estado)
}
