package colmena_raspberry

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetByColmenaUseCase struct {
	repo repositories.ColmenaRaspberryRepositories
}

func NewGetByColmenaUseCase(repo repositories.ColmenaRaspberryRepositories) *GetByColmenaUseCase {
	return &GetByColmenaUseCase{
		repo: repo,
	}
}

func (uc *GetByColmenaUseCase) Execute(id int) (entities.ColmenaRaspberryPi, error) {
	return uc.repo.GetColmenaRaspberry(id)
}
