package tipo_sensores

import (
	"apisensores/src/domain/entities"
	"apisensores/src/domain/repositories"
)

type UpdateTipoSensoresUseCase struct {
	repo repositories.TipoSensoresRepositores	
}

func NewUpdateTipoSensoresUseCase(repo repositories.TipoSensoresRepositores) *UpdateTipoSensoresUseCase {
	return &UpdateTipoSensoresUseCase{
		repo: repo,
	}
}

func (uc *UpdateTipoSensoresUseCase) Execute(tipoSensor entities.Tipo_Sesnores) error {
	return uc.repo.UpdateTipoSensor(tipoSensor)
}
