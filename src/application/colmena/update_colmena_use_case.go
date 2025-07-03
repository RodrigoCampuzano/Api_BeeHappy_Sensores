package colmena

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type UpdateColmenaUseCase struct {
	repo repositories.ColmenaRepositories
}

func NewUpdateColmenaUseCase(repo repositories.ColmenaRepositories) *UpdateColmenaUseCase {
	return &UpdateColmenaUseCase{
		repo: repo,
	}
}

func (uc *UpdateColmenaUseCase) Execute(colmena entities.Colmenas) error {
	return uc.repo.UpdateColmena(colmena)
}
