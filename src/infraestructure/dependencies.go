package infraestructure

import (
	"apisensores/src/application/calibracion"
	"apisensores/src/application/colmena"
	"apisensores/src/application/colmena_raspberry"
	"apisensores/src/application/raspberrypi"
	"apisensores/src/application/sensores"
	"apisensores/src/application/tipo_sensores"
	"apisensores/src/infraestructure/controllers"
	"apisensores/src/infraestructure/repositories/mysql"
	"apisensores/src/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func InitCalibracion(r gin.IRouter) {
	repo := mysql.NewCalibracionMySQL()
	createHandler := calibracion.NewCreateCalibracionUseCase(repo)
	getHandler := calibracion.NewGetCalibracionUseCase(repo)
	updateHandler := calibracion.NewUpdateCalibracionUseCase(repo)
	deleteHandler := calibracion.NewDeleteCalibracionUseCase(repo)

	controller := controllers.NewCalibracionController(createHandler, getHandler, updateHandler, deleteHandler)
	routes.CalibracionRoutes(r, controller)
}

func InitColmena(r gin.IRouter) {
	repo := mysql.NewColmenaMySQL()
	createHandler := colmena.NewCreateColmenaUseCase(repo)
	getHandler := colmena.NewGetColmenaUseCase(repo)
	updateHandler := colmena.NewUpdateColmenaUseCase(repo)
	deleteHandler := colmena.NewDeleteColmenaUseCase(repo)
	getAllHandler := colmena.NewGetAllColmenasUseCase(repo)
	updateEstadoHandler := colmena.NewUpdateEstadoColmenaUseCase(repo)

	controller := controllers.NewColmenaController(createHandler, getHandler, updateHandler, deleteHandler, getAllHandler, updateEstadoHandler)
	routes.ColmenaRoutes(r, controller)
}

func InitColmenaRaspberry(r gin.IRouter) {
	repo := mysql.NewColmenaRaspberryMySQL()
	createHandler := colmena_raspberry.NewCreateColmenaRaspberryUseCase(repo)
	getHandler := colmena_raspberry.NewGetByColmenaUseCase(repo)
	updateHandler := colmena_raspberry.NewUpdateColmenaRaspberryUseCase(repo)
	deleteHandler := colmena_raspberry.NewDeleteColmenaRaspberryUseCase(repo)
	controller := controllers.NewColmenaRaspberryController(createHandler, getHandler, updateHandler, deleteHandler,)
	routes.ColmenaRaspberryRoutes(r, controller)
}

func InitRaspberryPi(r gin.IRouter) {
	repo := mysql.NewRaspberryPiMySQL()
	createHandler := raspberrypi.NewCreateRaspberrypiUseCase(repo)
	getHandler := raspberrypi.NewGetRaspberryPiUseCase(repo)
	updateHandler := raspberrypi.NewUpdateRaspberryPiUseCase(repo)
	deleteHandler := raspberrypi.NewDeleteRaspberrypiUseCase(repo)
	getAllHandler := raspberrypi.NewGetAllRaspberryPiUseCase(repo)
	getByMACHandler := raspberrypi.NewGetByMACUseCase(repo)
	updateEstadoHandler := raspberrypi.NewUpdateEstadoRaspberryPiUseCase(repo)

	controller := controllers.NewRaspberrypiController(createHandler, getHandler, updateHandler, deleteHandler, getAllHandler, updateEstadoHandler, getByMACHandler)
	routes.RaspberryPiRoutes(r, controller)
}

func InitSensores(r gin.IRouter) {
	repo := mysql.NewSensoresMySQL()
	createHandler := sensores.NewCreateSensoresUseCase(repo)
	getHandler := sensores.NewGetSensoresUseCase(repo)
	updateHandler := sensores.NewUpdateSensoresUseCase(repo)
	deleteHandler := sensores.NewDeleteSensoresUseCase(repo)
	getAllHandler := sensores.NewGetAllSensoresUseCase(repo)
	getByRaspberryHandler := sensores.NewGetByRaspberryUseCase(repo)
	updateEstadoHandler := sensores.NewUpdateEstadoSensorUseCase(repo)

	controller := controllers.NewSensoresController(createHandler, getHandler, updateHandler, deleteHandler, getAllHandler, getByRaspberryHandler, updateEstadoHandler)
	routes.SensoresRoutes(r, controller)
}

func InitTipoSensores(r gin.IRouter) {
	repo := mysql.NewTipoSensoresMySQL()
	createHandler := tipo_sensores.NewCreateTipoSensoresUseCase(repo)
	getHandler := tipo_sensores.NewGetTipoSensoresUseCase(repo)
	updateHandler := tipo_sensores.NewUpdateTipoSensoresUseCase(repo)
	deleteHandler := tipo_sensores.NewDeleteTipoSensoresUseCase(repo)
	getAllHandler := tipo_sensores.NewGetAllTipoSensoresUseCase(repo)
	getByNombreHandler := tipo_sensores.NewGetByNombreUseCase(repo)

	controller := controllers.NewTipoSensoresController(createHandler, getHandler, updateHandler, deleteHandler, getAllHandler, getByNombreHandler)
	routes.TipoSensoresRoutes(r, controller)
}
