package tipo_sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetTipoSensoresUseCase struct {
	repo repositories.TipoSensoresRepositores	
}

func NewGetTipoSensoresUseCase(repo repositories.TipoSensoresRepositores) *GetTipoSensoresUseCase {
	return &GetTipoSensoresUseCase{
		repo: repo,
	}
}

func (uc *GetTipoSensoresUseCase) Execute(id int) (entities.Tipo_Sesnores, error) {
	return uc.repo.GetTipoSensor(id)
}
