package controllers

import (
	"apisensores/src/application/sensores"
	handlers "apisensores/src/infraestructure/handlers/sensores"

	"github.com/gin-gonic/gin"
)

type SensoresController struct {
	createSensoresHandler *handlers.CreateSensoresHandler
	getSensoresHandler    *handlers.GetSensoresHandler
	updateSensoresHandler *handlers.UpdateSensoresHandler
	deleteSensoresHandler *handlers.DeleteSensoresHandler
	getAllSensoresHandler *handlers.GetAllSensoresHandler
	getByRaspberryIDHandler *handlers.GetRaspidHandler
	updateEstadoSensoresHandler *handlers.UpdateEstadoSensorHandler	
}

func NewSensoresController(
	createSensoresUseCase *sensores.CreateSensoresUseCase,
	getSensoresUseCase *sensores.GetSensoresUseCase,
	updateSensoresUseCase *sensores.UpdateSensoresUseCase,
	deleteSensoresUseCase *sensores.DeleteSensoresUseCase,
	getAllSensoresUseCase *sensores.GetAllSensoresUseCase,
	updateEstadoSensoresUseCase *sensores.UpdateEstadoSensorUseCase,
	getByRaspberryIDUseCase *sensores.GetByRaspberryUseCase,
) *SensoresController {
	return &SensoresController{
		createSensoresHandler: handlers.NewCreateSensoresHandler(createSensoresUseCase),
		getSensoresHandler:    handlers.NewGetSensoresHandler(getSensoresUseCase),
		updateSensoresHandler: handlers.NewUpdateSensoresHandler(updateSensoresUseCase),
		deleteSensoresHandler: handlers.NewDeleteSensoresHandler(deleteSensoresUseCase),
		getAllSensoresHandler: handlers.NewGetAllSensoresHandler(getAllSensoresUseCase),
		getByRaspberryIDHandler: handlers.NewGetRaspidHandler(getByRaspberryIDUseCase),
		updateEstadoSensoresHandler: handlers.NewUpdateEstadoSensorHandler(updateEstadoSensoresUseCase),
			}
}

func (sc *SensoresController) CreateSensor(ctx *gin.Context) {
	sc.createSensoresHandler.Handle(ctx)
}

func (sc *SensoresController) GetSensor(ctx *gin.Context) {
	sc.getSensoresHandler.Handle(ctx)
}

func (sc *SensoresController) UpdateSensor(ctx *gin.Context) {
	sc.updateSensoresHandler.Handle(ctx)
}

func (sc *SensoresController) DeleteSensor(ctx *gin.Context) {
	sc.deleteSensoresHandler.Handle(ctx)
}

func (sc *SensoresController) GetAllSensores(ctx *gin.Context) {
	sc.getAllSensoresHandler.Handle(ctx)
}

func (sc *SensoresController) GetByRaspberryID(ctx *gin.Context) {
	sc.getByRaspberryIDHandler.Handle(ctx)
}

func (sc *SensoresController) UpdateEstadoSensores(ctx *gin.Context) {
	sc.updateEstadoSensoresHandler.Handle(ctx)
}