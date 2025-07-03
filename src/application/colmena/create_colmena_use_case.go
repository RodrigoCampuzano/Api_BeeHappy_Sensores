package colmena

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type CreateColmenaUseCase struct {
	repo repositories.ColmenaRepositories
}

func NewCreateColmenaUseCase(repo repositories.ColmenaRepositories) *CreateColmenaUseCase {
	return &CreateColmenaUseCase{
		repo: repo,
	}
}

func (uc *CreateColmenaUseCase) Execute(colmena entities.Colmenas) error {
	return uc.repo.CreateColmena(colmena)
}
