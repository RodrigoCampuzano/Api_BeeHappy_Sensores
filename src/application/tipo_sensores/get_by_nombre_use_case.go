package tipo_sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type GetByNombreUseCase struct {
	repo repositories.TipoSensoresRepositores			
}

func NewGetByNombreUseCase(repo repositories.TipoSensoresRepositores) *GetByNombreUseCase {
	return &GetByNombreUseCase{
		repo: repo,
	}
}

func (uc *GetByNombreUseCase) Execute(nombre string) (entities.Tipo_Sesnores, error) {
	return uc.repo.GetByNombre(nombre)
}
