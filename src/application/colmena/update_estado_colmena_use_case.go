package colmena

import (
	"apisensores/src/domain/repositories"
)

type UpdateEstadoColmenaUseCase struct {
	repo repositories.ColmenaRepositories
}

func NewUpdateEstadoColmenaUseCase(repo repositories.ColmenaRepositories) *UpdateEstadoColmenaUseCase {
	return &UpdateEstadoColmenaUseCase{
		repo: repo,
	}
}

func (uc *UpdateEstadoColmenaUseCase) Execute(id int, estado string) error {
	return uc.repo.UpdateEstadoColmena(id, estado)
}
