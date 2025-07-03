package repositories

import "apisensores/src/domain/entities"

type ColmenaRepositories interface {
	CreateColmena(colmena entities.Colmenas) error
	GetColmena(id int) (entities.Colmenas, error)
	UpdateColmena(colmena entities.Colmenas) error
	DeleteColmena(id int) error
	GetAllColmenas() ([]entities.Colmenas, error)
	UpdateEstadoColmena(id int, estado string) error
}
