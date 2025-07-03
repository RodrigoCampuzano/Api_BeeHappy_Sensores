package controllers

import (
	"apisensores/src/application/tipo_sensores"
	handlers "apisensores/src/infraestructure/handlers/tipo_sensores"

	"github.com/gin-gonic/gin"
)

type TipoSensoresController struct {
	createTipoSensoresHandler *handlers.CreateTipoSensoresHandler
	getTipoSensoresHandler    *handlers.GetTipoSensoresHandler
	updateTipoSensoresHandler *handlers.UpdateTipoSensoresHandler
	deleteTipoSensoresHandler *handlers.DeleteTipoSensoresHandler
	getAllTipoSensoresHandler *handlers.GetAllTipoSensoresHandler
	getByNombreHandler *handlers.GetByNombreTipoSensoresHandler
}

func NewTipoSensoresController(
	createTipoSensoresUseCase *tipo_sensores.CreateTipoSensoresUseCase,
	getTipoSensoresUseCase *tipo_sensores.GetTipoSensoresUseCase,
	updateTipoSensoresUseCase *tipo_sensores.UpdateTipoSensoresUseCase,
	deleteTipoSensoresUseCase *tipo_sensores.DeleteTipoSensoresUseCase,
	getAllTipoSensoresUseCase *tipo_sensores.GetAllTipoSensoresUseCase,
	getByNombreUseCase *tipo_sensores.GetByNombreUseCase,
) *TipoSensoresController {
	return &TipoSensoresController{
		createTipoSensoresHandler: handlers.NewCreateTipoSensoresHandler(createTipoSensoresUseCase),
		getTipoSensoresHandler:    handlers.NewGetTipoSensoresHandler(getTipoSensoresUseCase),
		updateTipoSensoresHandler: handlers.NewUpdateTipoSensoresHandler(updateTipoSensoresUseCase),
		deleteTipoSensoresHandler: handlers.NewDeleteTipoSensoresHandler(deleteTipoSensoresUseCase),
		getAllTipoSensoresHandler: handlers.NewGetAllTipoSensoresHandler(getAllTipoSensoresUseCase),
		getByNombreHandler: handlers.NewGetByNombreTipoSensoresHandler(getByNombreUseCase),
	}
}

func (tsc *TipoSensoresController) CreateTipoSensor(ctx *gin.Context) {
	tsc.createTipoSensoresHandler.Handle(ctx)
}

func (tsc *TipoSensoresController) GetTipoSensor(ctx *gin.Context) {
	tsc.getTipoSensoresHandler.Handle(ctx)
}

func (tsc *TipoSensoresController) UpdateTipoSensor(ctx *gin.Context) {
	tsc.updateTipoSensoresHandler.Handle(ctx)
}

func (tsc *TipoSensoresController) DeleteTipoSensor(ctx *gin.Context) {
	tsc.deleteTipoSensoresHandler.Handle(ctx)
}

func (tsc *TipoSensoresController) GetAllTipoSensores(ctx *gin.Context) {
	tsc.getAllTipoSensoresHandler.Handle(ctx)
}

func (tsc *TipoSensoresController) GetByNombre(ctx *gin.Context) {
	tsc.getByNombreHandler.Handle(ctx)
}