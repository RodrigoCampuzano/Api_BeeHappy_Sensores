package tipo_sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type CreateTipoSensoresUseCase struct {
	repo repositories.TipoSensoresRepositores		
}

func NewCreateTipoSensoresUseCase(repo repositories.TipoSensoresRepositores) *CreateTipoSensoresUseCase {
	return &CreateTipoSensoresUseCase{
		repo: repo,
	}
}

func (uc *CreateTipoSensoresUseCase) Execute(tipoSensor entities.Tipo_Sesnores) error {
	return uc.repo.CreateTipoSensor(tipoSensor)
}
