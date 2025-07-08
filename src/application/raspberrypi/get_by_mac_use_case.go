package raspberrypi

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetByMACUseCase struct {
	repo repositories.RaspberryPiRepositories
}

func NewGetByMACUseCase(repo repositories.RaspberryPiRepositories) *GetByMACUseCase {
	return &GetByMACUseCase{
		repo: repo,	
	}
}

func (uc *GetByMACUseCase) Execute(mac string) (entities.Raspberrypi, error) {
	return uc.repo.GetByMAC(mac)
}
