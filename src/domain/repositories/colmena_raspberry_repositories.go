package repositories

import "apisensores/src/domain/entities"

type ColmenaRaspberryRepositories interface {
	CreateColmenaRaspberry(cr entities.ColmenaRaspberryPi) error
	GetColmenaRaspberry(id int) (entities.ColmenaRaspberryPi, error)
	UpdateColmenaRaspberry(cr entities.ColmenaRaspberryPi) error
	DeleteColmenaRaspberry(id int) error
}
