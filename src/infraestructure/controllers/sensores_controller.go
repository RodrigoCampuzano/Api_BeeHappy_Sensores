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
	getByRaspberryHandler *handlers.GetRaspidHandler
	updateEstadoSensoresHandler *handlers.UpdateEstadoSensorHandler	
}

func NewSensoresController(
	createSensoresUseCase *sensores.CreateSensoresUseCase,
	getSensoresUseCase *sensores.GetSensoresUseCase,
	updateSensoresUseCase *sensores.UpdateSensoresUseCase,
	deleteSensoresUseCase *sensores.DeleteSensoresUseCase,
	getAllSensoresUseCase *sensores.GetAllSensoresUseCase,
	updateEstadoSensoresUseCase *sensores.UpdateEstadoSensorUseCase,
	getByRaspberryUseCase *sensores.GetByRaspberryUseCase,
) *SensoresController {
	return &SensoresController{
		createSensoresHandler: handlers.NewCreateSensoresHandler(createSensoresUseCase),
		getSensoresHandler:    handlers.NewGetSensoresHandler(getSensoresUseCase),
		updateSensoresHandler: handlers.NewUpdateSensoresHandler(updateSensoresUseCase),
		deleteSensoresHandler: handlers.NewDeleteSensoresHandler(deleteSensoresUseCase),
		getAllSensoresHandler: handlers.NewGetAllSensoresHandler(getAllSensoresUseCase),
		getByRaspberryHandler: handlers.NewGetRaspidHandler(getByRaspberryUseCase),
		updateEstadoSensoresHandler: handlers.NewUpdateEstadoSensorHandler(updateEstadoSensoresUseCase),
			}
}

//@Summary Crear un sensor
//@Description Crea un nuevo sensor en la base de datos
//@Accept json
//@Produce json
//@Param sensor body entities.CreateSensorRequest true "Sensor a crear"
//@Success 200 {object} entities.SensorResponse
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores [post]	
func (sc *SensoresController) CreateSensor(ctx *gin.Context) {
	sc.createSensoresHandler.Handle(ctx)
}

//@Summary Obtener un sensor
//@Description Obtiene un sensor por su ID
//@Accept json
//@Produce json
//@Param id path int true "ID del sensor"
//@Success 200 {object} entities.SensorResponse
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores/{id} [get]
func (sc *SensoresController) GetSensor(ctx *gin.Context) {
	sc.getSensoresHandler.Handle(ctx)
}

//@Summary Actualizar un sensor
//@Description Actualiza un sensor existente
//@Accept json
//@Produce json
//@Param id path int true "ID del sensor"
//@Param sensor body entities.UpdateSensorRequest true "Sensor a actualizar"
//@Success 200 {object} entities.SensorResponse
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores/{id} [put]
func (sc *SensoresController) UpdateSensor(ctx *gin.Context) {
	sc.updateSensoresHandler.Handle(ctx)
}

//@Summary Eliminar un sensor
//@Description Elimina un sensor existente
//@Accept json
//@Produce json
//@Param id path int true "ID del sensor"
//@Success 200 {object} map[string]string
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores/{id} [delete]
func (sc *SensoresController) DeleteSensor(ctx *gin.Context) {
	sc.deleteSensoresHandler.Handle(ctx)
}

//@Summary Obtener todos los sensores
//@Description Obtiene todos los sensores registrados
//@Accept json
//@Produce json
//@Success 200 {array} entities.SensorResponse
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores [get]
func (sc *SensoresController) GetAllSensores(ctx *gin.Context) {
	sc.getAllSensoresHandler.Handle(ctx)
}

//@Summary Obtener sensores por Raspberry Pi
//@Description Obtiene todos los sensores asociados a un Raspberry Pi
//@Accept json
//@Produce json
//@Param id path int true "ID del Raspberry Pi"
//@Success 200 {array} entities.SensorResponse
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores/raspberry/{id} [get]
func (sc *SensoresController) GetByRaspberry(ctx *gin.Context) {
	sc.getByRaspberryHandler.Handle(ctx)
}

//@Summary Actualizar el estado de un sensor
//@Description Actualiza el estado de un sensor existente
//@Accept json
//@Produce json
//@Param id path int true "ID del sensor"
//@Param estado body entities.UpdateEstadoSensorRequest true "Estado a actualizar"
//@Success 200 {object} map[string]string
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /sensores/{id}/estado [put]
func (sc *SensoresController) UpdateEstadoSensores(ctx *gin.Context) {
	sc.updateEstadoSensoresHandler.Handle(ctx)
}