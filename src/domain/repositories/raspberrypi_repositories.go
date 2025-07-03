package repositories

import "apisensores/src/domain/entities"

type RaspberryPiRepositories interface {
	CreateRaspberryPi(rpi entities.Raspberrypi) error
	GetRaspberryPi(id int) (entities.Raspberrypi, error)
	UpdateRaspberryPi(rpi entities.Raspberrypi) error
	DeleteRaspberryPi(id int) error
	GetAllRaspberryPi() ([]entities.Raspberrypi, error)
	UpdateEstadoRaspberryPi(id int, estado string) error
	GetByMAC(mac string) ([]*entities.Raspberrypi, error)
}
