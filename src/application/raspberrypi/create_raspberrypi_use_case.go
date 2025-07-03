package raspberrypi

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type CreateRaspberrypiUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewCreateRaspberrypiUseCase(repo repositories.RaspberryPiRepositories) *CreateRaspberrypiUseCase {
	return &CreateRaspberrypiUseCase{
		repo: repo,
	}
}

func (uc *CreateRaspberrypiUseCase) Execute(raspberrypi entities.Raspberrypi) error {
	return uc.repo.CreateRaspberryPi(raspberrypi)
}
