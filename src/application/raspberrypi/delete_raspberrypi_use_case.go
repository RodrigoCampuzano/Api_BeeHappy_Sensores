package raspberrypi

import (
	"apisensores/src/domain/repositories"
)

type DeleteRaspberrypiUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewDeleteRaspberrypiUseCase(repo repositories.RaspberryPiRepositories) *DeleteRaspberrypiUseCase {
	return &DeleteRaspberrypiUseCase{
		repo: repo,
	}
}

func (uc *DeleteRaspberrypiUseCase) Execute(id int) error {
	return uc.repo.DeleteRaspberryPi(id)
}
