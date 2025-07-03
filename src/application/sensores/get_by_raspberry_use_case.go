package sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetByRaspberryUseCase struct {
	repo repositories.SensoresRepositories	
}

func NewGetByRaspberryUseCase(repo repositories.SensoresRepositories) *GetByRaspberryUseCase {
	return &GetByRaspberryUseCase{
		repo: repo,
	}
}

func (uc *GetByRaspberryUseCase) Execute(raspberryID int) ([]entities.Sensores, error) {
	return uc.repo.GetByRaspberryID(raspberryID)
}
