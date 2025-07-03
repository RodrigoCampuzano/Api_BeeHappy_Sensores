package colmena

import (
	"apisensores/src/domain/repositories"
)

type DeleteColmenaUseCase struct {
	repo repositories.ColmenaRepositories
}

func NewDeleteColmenaUseCase(repo repositories.ColmenaRepositories) *DeleteColmenaUseCase {
	return &DeleteColmenaUseCase{
		repo: repo,
	}
}

func (uc *DeleteColmenaUseCase) Execute(id int) error {
	return uc.repo.DeleteColmena(id)
}
