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

//@Summary Crear un nuevo tipo de sensor
//@Description Crea un nuevo tipo de sensor en la base de datos
//@Tags Tipo Sensores
//@Accept json
//@Produce json
//@Param TipoSensor body entities.CreateTipoSensorRequest true "Tipo de sensor a crear"
//@Success 201 {object} entities.TipoSensorResponse "Tipo de sensor creado exitosamente"	
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /tipos-sensores [post]
func (tsc *TipoSensoresController) CreateTipoSensor(ctx *gin.Context) {
	tsc.createTipoSensoresHandler.Handle(ctx)
}

//@Summary Obtener un tipo de sensor por su ID
//@Description Obtiene un tipo de sensor por su ID de la base de datos
//@Tags Tipo Sensores
//@Accept json
//@Produce json
//@Param id path int true "ID del tipo de sensor"
//@Success 200 {object} entities.TipoSensorResponse "Tipo de sensor obtenido exitosamente"
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /tipos-sensores/{id} [get]
func (tsc *TipoSensoresController) GetTipoSensor(ctx *gin.Context) {
	tsc.getTipoSensoresHandler.Handle(ctx)
}

//@Summary Actualizar un tipo de sensor existente
//@Description Actualiza un tipo de sensor existente en la base de datos
//@Tags Tipo Sensores
//@Accept json
//@Produce json
//@Param id path int true "ID del tipo de sensor"
//@Param TipoSensor body entities.UpdateTipoSensorRequest true "Tipo de sensor a actualizar"
//@Success 200 {object} entities.TipoSensorResponse "Tipo de sensor actualizado exitosamente"
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /tipos-sensores/{id} [put]
func (tsc *TipoSensoresController) UpdateTipoSensor(ctx *gin.Context) {
	tsc.updateTipoSensoresHandler.Handle(ctx)
}

//@Summary Eliminar un tipo de sensor existente
//@Description Elimina un tipo de sensor existente de la base de datos
//@Tags Tipo Sensores
//@Accept json
//@Produce json
//@Param id path int true "ID del tipo de sensor"
//@Success 200 {object} entities.TipoSensorResponse "Tipo de sensor eliminado exitosamente"
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /tipos-sensores/{id} [delete]
func (tsc *TipoSensoresController) DeleteTipoSensor(ctx *gin.Context) {
	tsc.deleteTipoSensoresHandler.Handle(ctx)
}

//@Summary Obtener todos los tipos de sensores
//@Description Obtiene todos los tipos de sensores de la base de datos
//@Tags Tipo Sensores
//@Accept json
//@Produce json
//@Success 200 {array} entities.TipoSensorResponse "Lista de tipos de sensores obtenida exitosamente"
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /tipos-sensores [get]
func (tsc *TipoSensoresController) GetAllTipoSensores(ctx *gin.Context) {
	tsc.getAllTipoSensoresHandler.Handle(ctx)
}

//@Summary Obtener un tipo de sensor por su nombre
//@Description Obtiene un tipo de sensor por su nombre de la base de datos
//@Tags Tipo Sensores
//@Accept json
//@Produce json
//@Param nombre path string true "Nombre del tipo de sensor"
//@Success 200 {object} entities.TipoSensorResponse "Tipo de sensor obtenido exitosamente"
//@Failure 400 {object} map[string]string
//@Failure 500 {object} map[string]string
//@Security ApiKeyAuth
//@Router /tipos-sensores/nombre/{nombre} [get]
func (tsc *TipoSensoresController) GetByNombre(ctx *gin.Context) {
	tsc.getByNombreHandler.Handle(ctx)
}