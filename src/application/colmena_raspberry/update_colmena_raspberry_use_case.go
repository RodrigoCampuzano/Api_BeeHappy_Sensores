package colmena_raspberry

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type UpdateColmenaRaspberryUseCase struct {
	repo repositories.ColmenaRaspberryRepositories
}

func NewUpdateColmenaRaspberryUseCase(repo repositories.ColmenaRaspberryRepositories) *UpdateColmenaRaspberryUseCase {
	return &UpdateColmenaRaspberryUseCase{
		repo: repo,
	}
}

func (uc *UpdateColmenaRaspberryUseCase) Execute(colmenaRaspberry entities.ColmenaRaspberryPi) error {
	return uc.repo.UpdateColmenaRaspberry(colmenaRaspberry)
}
