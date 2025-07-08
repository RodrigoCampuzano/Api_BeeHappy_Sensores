package colmena_raspberry

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetByColmenaRaspberryUseCase struct {
	repo repositories.ColmenaRaspberryRepositories
}

func NewGetByColmenaRaspberryUseCase(repo repositories.ColmenaRaspberryRepositories) *GetByColmenaRaspberryUseCase {
	return &GetByColmenaRaspberryUseCase{
		repo: repo,
	}
}

func (uc *GetByColmenaRaspberryUseCase) Execute(id int) (entities.ColmenaRaspberryPi, error) {
	return uc.repo.GetColmenaRaspberry(id)
}
