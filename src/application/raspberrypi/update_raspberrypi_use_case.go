package raspberrypi

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type UpdateRaspberryPiUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewUpdateRaspberryPiUseCase(repo repositories.RaspberryPiRepositories) *UpdateRaspberryPiUseCase {
	return &UpdateRaspberryPiUseCase{
		repo: repo,
	}
}

func (uc *UpdateRaspberryPiUseCase) Execute(id int, raspberrypi entities.Raspberrypi) error {
	return uc.repo.UpdateRaspberryPi(id, raspberrypi)
}
