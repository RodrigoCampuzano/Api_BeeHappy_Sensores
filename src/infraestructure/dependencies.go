package infraestructure

import (
	"apisensores/src/application/calibracion"
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
