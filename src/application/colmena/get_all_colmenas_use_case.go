package colmena

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetAllColmenasUseCase struct {
	repo repositories.ColmenaRepositories
}

func NewGetAllColmenasUseCase(repo repositories.ColmenaRepositories) *GetAllColmenasUseCase {
	return &GetAllColmenasUseCase{
		repo: repo,
	}
}

func (uc *GetAllColmenasUseCase) Execute() ([]entities.Colmenas, error) {
	return uc.repo.GetAllColmenas()
}
