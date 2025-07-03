package raspberrypi

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetRaspberryPiUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewGetRaspberryPiUseCase(repo repositories.RaspberryPiRepositories) *GetRaspberryPiUseCase {
	return &GetRaspberryPiUseCase{
		repo: repo,
	}
}

func (uc *GetRaspberryPiUseCase) Execute(id int) ([]entities.Raspberrypi, error) {
	return uc.repo.GetRaspberryPi(id)
}
