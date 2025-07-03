package raspberrypi

import (
	"apisensores/src/domain/repositories"
)

type UpdateEstadoRaspberryPiUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewUpdateEstadoRaspberryPiUseCase(repo repositories.RaspberryPiRepositories) *UpdateEstadoRaspberryPiUseCase {
	return &UpdateEstadoRaspberryPiUseCase{
		repo: repo,
	}
}

func (uc *UpdateEstadoRaspberryPiUseCase) Execute(id int, estado string) error {
	return uc.repo.UpdateEstadoRaspberryPi(id, estado)
}
