package repositories

import "apisensores/src/domain/entities"

type TipoSensoresRepositores interface {
	CreateTipoSensor(tipoSensor entities.Tipo_Sesnores) error
	GetTipoSensor(id int) (entities.Tipo_Sesnores, error)
	UpdateTipoSensor(tipoSensor entities.Tipo_Sesnores) error
	DeleteTipoSensor(id int) error
	GetAllTipoSensores() ([]entities.Tipo_Sesnores, error)
	GetByNombre(nombre string) (entities.Tipo_Sesnores, error)
}
