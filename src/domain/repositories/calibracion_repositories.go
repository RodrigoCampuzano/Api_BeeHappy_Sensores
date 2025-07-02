package repositories

import "apisensores/src/domain/entities"

type CalibracionRepositories interface {
	CreateCalibracion(calibracion entities.Calibracion) error
	GetCalibracion(id int) (entities.Calibracion, error)
	UpdateCalibracion(calibracion entities.Calibracion) error
	DeleteCalibracion(id int) error
}
