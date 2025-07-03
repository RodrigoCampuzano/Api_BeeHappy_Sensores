package colmena_raspberry

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type CreateColmenaRaspberryUseCase struct {
	repo repositories.ColmenaRaspberryRepositories
}

func NewCreateColmenaRaspberryUseCase(repo repositories.ColmenaRaspberryRepositories) *CreateColmenaRaspberryUseCase {
	return &CreateColmenaRaspberryUseCase{
		repo: repo,
	}
}

func (uc *CreateColmenaRaspberryUseCase) Execute(colmenaRaspberry entities.ColmenaRaspberryPi) error {
	return uc.repo.CreateColmenaRaspberry(colmenaRaspberry)
}
