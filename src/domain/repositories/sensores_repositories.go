package repositories

import "apisensores/src/domain/entities"

type SensoresRepositories interface {
	CreateSensor(sensor entities.Sensores) error
	GetSensor(id int) (entities.Sensores, error)
	UpdateSensor(sensor entities.Sensores) error
	DeleteSensor(id int) error
	GetAllSensores() ([]entities.Sensores, error)
	GetByRaspberryID(raspberryID int) ([]entities.Sensores, error)
	UpdateEstadoSensor(id int, estado string) error
}
