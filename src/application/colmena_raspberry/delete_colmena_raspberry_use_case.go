package colmena_raspberry

import (
	"apisensores/src/domain/repositories"
)

type DeleteColmenaRaspberryUseCase struct {
	repo repositories.ColmenaRaspberryRepositories
}

func NewDeleteColmenaRaspberryUseCase(repo repositories.ColmenaRaspberryRepositories) *DeleteColmenaRaspberryUseCase {
	return &DeleteColmenaRaspberryUseCase{
		repo: repo,
	}
}

func (uc *DeleteColmenaRaspberryUseCase) Execute(id int) error {
	return uc.repo.DeleteColmenaRaspberry(id)
} 