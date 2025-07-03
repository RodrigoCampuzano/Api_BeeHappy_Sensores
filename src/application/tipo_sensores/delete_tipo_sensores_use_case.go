package tipo_sensores

import (
	"apisensores/src/domain/repositories"
)

type DeleteTipoSensoresUseCase struct {
	repo repositories.TipoSensoresRepositores
}

func NewDeleteTipoSensoresUseCase(repo repositories.TipoSensoresRepositores) *DeleteTipoSensoresUseCase {
	return &DeleteTipoSensoresUseCase{
		repo: repo,
	}
}

func (uc *DeleteTipoSensoresUseCase) Execute(id int) error {
	return uc.repo.DeleteTipoSensor(id)
}
