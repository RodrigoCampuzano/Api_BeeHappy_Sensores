package raspberrypi

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetAllRaspberryPiUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewGetAllRaspberryPiUseCase(repo repositories.RaspberryPiRepositories) *GetAllRaspberryPiUseCase {
	return &GetAllRaspberryPiUseCase{
		repo: repo,
	}
}

func (uc *GetAllRaspberryPiUseCase) Execute() ([]entities.Raspberrypi, error) {
	return uc.repo.GetAllRaspberryPi()
}
