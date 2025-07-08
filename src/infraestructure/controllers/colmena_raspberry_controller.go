package controllers

import (
	"apisensores/src/application/colmena_raspberry"
	handlers "apisensores/src/infraestructure/handlers/colmena_raspberry"

	"github.com/gin-gonic/gin"
)

type ColmenaRaspberryController struct {
	createColmenaRaspberryHandler *handlers.CreateColmenaRaspberryHandler
	getByColmenaHandler           *handlers.GetColmenaRaspberryHandler
	updateColmenaRaspberryHandler *handlers.UpdateColmenaRaspberryHandler
	deleteColmenaRaspberryHandler *handlers.DeleteColmenaRaspberryHandler
}

// Constructor
func NewColmenaRaspberryController(
	createUseCase *colmena_raspberry.CreateColmenaRaspberryUseCase,
	getByColmenaUseCase *colmena_raspberry.GetByColmenaRaspberryUseCase,
	updateUseCase *colmena_raspberry.UpdateColmenaRaspberryUseCase,
	deleteUseCase *colmena_raspberry.DeleteColmenaRaspberryUseCase,
) *ColmenaRaspberryController {
	return &ColmenaRaspberryController{
		createColmenaRaspberryHandler: handlers.NewCreateColmenaRaspberryHandler(createUseCase),
		getByColmenaHandler:           handlers.NewGetColmenaRaspberryHandler(getByColmenaUseCase),
		updateColmenaRaspberryHandler: handlers.NewUpdateColmenaRaspberryHandler(updateUseCase),
		deleteColmenaRaspberryHandler: handlers.NewDeleteColmenaRaspberryHandler(deleteUseCase),
	}
}

// @Summary Crear una nueva relación colmena-raspberry
// @Description Crea una nueva relación entre una colmena y un dispositivo Raspberry Pi
// @Tags Colmena-Raspberry
// @Accept json
// @Produce json
// @Param request body entities.CreateColmenaRaspberryRequest true "Datos de la relación"
// @Success 201 {object} entities.ColmenaRaspberryResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena-raspberry [post]
func (cr *ColmenaRaspberryController) CreateColmenaRaspberry(ctx *gin.Context) {
	cr.createColmenaRaspberryHandler.Handle(ctx)
}

// @Summary Obtener relación colmena-raspberry por ID
// @Description Obtiene los detalles de una relación colmena-raspberry por su ID
// @Tags Colmena-Raspberry
// @Accept json
// @Produce json
// @Param id path int true "ID de la relación"
// @Success 200 {object} entities.ColmenaRaspberryResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena-raspberry/{id} [get]
func (cr *ColmenaRaspberryController) GetColmenaRaspberry(ctx *gin.Context) {
	cr.getByColmenaHandler.Handle(ctx)
}

// @Summary Actualizar una relación colmena-raspberry
// @Description Actualiza una relación existente entre colmena y Raspberry Pi
// @Tags Colmena-Raspberry
// @Accept json
// @Produce json
// @Param id path int true "ID de la relación"
// @Param request body entities.UpdateColmenaRaspberryRequest true "Datos de actualización"
// @Success 200 {object} entities.ColmenaRaspberryResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena-raspberry/{id} [put]
func (cr *ColmenaRaspberryController) UpdateColmenaRaspberry(ctx *gin.Context) {
	cr.updateColmenaRaspberryHandler.Handle(ctx)
}

// @Summary Eliminar una relación colmena-raspberry
// @Description Elimina una relación existente entre colmena y Raspberry Pi
// @Tags Colmena-Raspberry
// @Accept json
// @Produce json
// @Param id path int true "ID de la relación"
// @Success 204 "No Content"
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena-raspberry/{id} [delete]
func (cr *ColmenaRaspberryController) DeleteColmenaRaspberry(ctx *gin.Context) {
	cr.deleteColmenaRaspberryHandler.Handle(ctx)
}
