package controllers

import (
	"apisensores/src/application/raspberrypi"
	handlers "apisensores/src/infraestructure/handlers/raspberrypi"

	"github.com/gin-gonic/gin"
)

type RaspberrypiController struct {
	createRaspberrypiHandler *handlers.CreateRaspberrypiHandler
	getRaspberrypiHandler    *handlers.GetRaspberrypiHandler
	updateRaspberrypiHandler *handlers.UpdateRaspberrypiHandler
	deleteRaspberrypiHandler *handlers.DeleteRaspberrypiHandler
	getAllRaspberrypiHandler *handlers.GetAllRaspberrypiHandler
	updateEstadoRaspberrypiHandler *handlers.UpdateEstadoRaspberrypiHandler
	getByMACHandler *handlers.GetMacRaspberrypiHandler
}

func NewRaspberrypiController(
	createRaspberrypiUseCase *raspberrypi.CreateRaspberrypiUseCase,
	getRaspberrypiUseCase *raspberrypi.GetRaspberryPiUseCase,
	updateRaspberrypiUseCase *raspberrypi.UpdateRaspberryPiUseCase,
	deleteRaspberrypiUseCase *raspberrypi.DeleteRaspberrypiUseCase,
	getAllRaspberrypiUseCase *raspberrypi.GetAllRaspberryPiUseCase,
	updateEstadoRaspberrypiUseCase *raspberrypi.UpdateEstadoRaspberryPiUseCase,
	getByMACUseCase *raspberrypi.GetByMACUseCase,
) *RaspberrypiController {
	return &RaspberrypiController{
		createRaspberrypiHandler: handlers.NewCreateRaspberrypiHandler(createRaspberrypiUseCase),
		getRaspberrypiHandler:    handlers.NewGetRaspberrypiHandler(getRaspberrypiUseCase),
		updateRaspberrypiHandler: handlers.NewUpdateRaspberrypiHandler(updateRaspberrypiUseCase),
		deleteRaspberrypiHandler: handlers.NewDeleteRaspberrypiHandler(deleteRaspberrypiUseCase),
		getAllRaspberrypiHandler: handlers.NewGetAllRaspberrypiHandler(getAllRaspberrypiUseCase),
		updateEstadoRaspberrypiHandler: handlers.NewUpdateEstadoRaspberrypiHandler(updateEstadoRaspberrypiUseCase),
		getByMACHandler: handlers.NewGetMacRaspberrypiHandler(getByMACUseCase),
	}
}

// @Summary Crear un nuevo Raspberry Pi
// @Description Registra un nuevo dispositivo Raspberry Pi en el sistema
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Param request body entities.CreateRaspberryPiRequest true "Datos del dispositivo"
// @Success 201 {object} entities.RaspberryPiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry [post]
func (rc *RaspberrypiController) CreateRaspberryPi(ctx *gin.Context) {
	rc.createRaspberrypiHandler.Handle(ctx)
}

// @Summary Obtener un Raspberry Pi
// @Description Obtiene los detalles de un dispositivo Raspberry Pi por su ID
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Success 200 {object} entities.RaspberryPiResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry/{id} [get]
func (rc *RaspberrypiController) GetRaspberryPi(ctx *gin.Context) {
	rc.getRaspberrypiHandler.Handle(ctx)
}

// @Summary Obtener todos los Raspberry Pi
// @Description Obtiene la lista de todos los dispositivos Raspberry Pi
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Success 200 {object} entities.RaspberryPiListResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry [get]
func (rc *RaspberrypiController) GetAllRaspberryPi(ctx *gin.Context) {
	rc.getAllRaspberrypiHandler.Handle(ctx)
}

// @Summary Obtener Raspberry Pi por MAC
// @Description Obtiene un dispositivo Raspberry Pi por su dirección MAC
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Param mac path string true "Dirección MAC del dispositivo"
// @Success 200 {object} entities.RaspberryPiResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry/mac/{mac} [get]
func (rc *RaspberrypiController) GetByMAC(ctx *gin.Context) {
	rc.getByMACHandler.Handle(ctx)
}

// @Summary Actualizar un Raspberry Pi
// @Description Actualiza los datos de un dispositivo Raspberry Pi existente
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Param request body entities.UpdateRaspberryPiRequest true "Datos de actualización"
// @Success 200 {object} entities.RaspberryPiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry/{id} [put]
func (rc *RaspberrypiController) UpdateRaspberryPi(ctx *gin.Context) {
	rc.updateRaspberrypiHandler.Handle(ctx)
}

// @Summary Actualizar estado de Raspberry Pi
// @Description Actualiza el estado de un dispositivo Raspberry Pi
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Param request body entities.UpdateEstadoRaspberryPiRequest true "Nuevo estado"
// @Success 200 {object} entities.RaspberryPiResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry/{id}/estado [put]
func (rc *RaspberrypiController) UpdateEstadoRaspberryPi(ctx *gin.Context) {
	rc.updateEstadoRaspberrypiHandler.Handle(ctx)
}

// @Summary Eliminar un Raspberry Pi
// @Description Elimina un dispositivo Raspberry Pi por su ID
// @Tags RaspberryPi
// @Accept json
// @Produce json
// @Param id path int true "ID del dispositivo"
// @Success 204 "No Content"
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /raspberry/{id} [delete]
func (rc *RaspberrypiController) DeleteRaspberryPi(ctx *gin.Context) {
	rc.deleteRaspberrypiHandler.Handle(ctx)
}