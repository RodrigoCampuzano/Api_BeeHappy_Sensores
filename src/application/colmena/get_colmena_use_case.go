package colmena

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetColmenaUseCase struct {
	repo repositories.ColmenaRepositories
}

func NewGetColmenaUseCase(repo repositories.ColmenaRepositories) *GetColmenaUseCase {
	return &GetColmenaUseCase{
		repo: repo,
	}
}

func (uc *GetColmenaUseCase) Execute(id int) (entities.Colmenas, error) {
	return uc.repo.GetColmena(id)
}
