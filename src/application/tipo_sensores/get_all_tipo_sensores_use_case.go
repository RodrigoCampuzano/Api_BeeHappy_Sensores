package tipo_sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetAllTipoSensoresUseCase struct {
	repo repositories.TipoSensoresRepositores
}

func NewGetAllTipoSensoresUseCase(repo repositories.TipoSensoresRepositores) *GetAllTipoSensoresUseCase {
	return &GetAllTipoSensoresUseCase{
		repo: repo,
	}
}

func (uc *GetAllTipoSensoresUseCase) Execute() ([]entities.Tipo_Sesnores, error) {
	return uc.repo.GetAllTipoSensores()
}
