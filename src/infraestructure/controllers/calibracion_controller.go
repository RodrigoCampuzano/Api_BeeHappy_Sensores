package controllers

import (
	"apisensores/src/application/calibracion"
	handlers "apisensores/src/infraestructure/handlers/calibracion"

	"github.com/gin-gonic/gin"
)

// CalibracionController estructura del controlador
type CalibracionController struct {
	createCalibracionHandler *handlers.CreateCalibracionHandler
	getCalibracionHandler    *handlers.GetCalibracionHandler
	updateCalibracionHandler *handlers.UpdateCalibracionHandler
	deleteCalibracionHandler *handlers.DeleteCalibracionHandler
}

func NewCalibracionController(
	createCalibracionUseCase *calibracion.CreateCalibracionUseCase,
	getCalibracionUseCase *calibracion.GetCalibracionUseCase,
	updateCalibracionUseCase *calibracion.UpdateCalibracionUseCase,
	deleteCalibracionUseCase *calibracion.DeleteCalibracionUseCase,
) *CalibracionController {
	return &CalibracionController{
		createCalibracionHandler: handlers.NewCreateCalibracionHandler(createCalibracionUseCase),
		getCalibracionHandler:    handlers.NewGetCalibracionHandler(getCalibracionUseCase),
		updateCalibracionHandler: handlers.NewUpdateCalibracionHandler(updateCalibracionUseCase),
		deleteCalibracionHandler: handlers.NewDeleteCalibracionHandler(deleteCalibracionUseCase),
	}
}

// @Summary Crear una nueva calibración
// @Description Crea un nuevo registro de calibración para un sensor
// @Tags Calibración
// @Accept json
// @Produce json
// @Param request body entities.CreateCalibracionRequest true "Datos de la calibración"
// @Success 201 {object} entities.CalibracionResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 401 {object} entities.ErrorResponse "No autorizado"
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /calibracion [post]
func (cc *CalibracionController) CreateCalibracion(ctx *gin.Context) {
	cc.createCalibracionHandler.Handle(ctx)
}

// @Summary Obtener una calibración
// @Description Obtiene los detalles de una calibración por su ID
// @Tags Calibración
// @Accept json
// @Produce json
// @Param id path int true "ID de la calibración"
// @Success 200 {object} entities.CalibracionResponse
// @Failure 401 {object} entities.ErrorResponse "No autorizado"
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /calibracion/{id} [get]
func (cc *CalibracionController) GetCalibracion(ctx *gin.Context) {
	cc.getCalibracionHandler.Handle(ctx)
}

// @Summary Actualizar una calibración
// @Description Actualiza los datos de una calibración existente
// @Tags Calibración
// @Accept json
// @Produce json
// @Param id path int true "ID de la calibración"
// @Param request body entities.UpdateCalibracionRequest true "Datos de actualización"
// @Success 200 {object} entities.CalibracionResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /calibracion/{id} [put]
func (cc *CalibracionController) UpdateCalibracion(ctx *gin.Context) {
	cc.updateCalibracionHandler.Handle(ctx)
}

// @Summary Eliminar una calibración
// @Description Elimina una calibración por su ID
// @Tags Calibración
// @Accept json
// @Produce json
// @Param id path int true "ID de la calibración"
// @Success 204 "No Content"
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /calibracion/{id} [delete]
func (cc *CalibracionController) DeleteCalibracion(ctx *gin.Context) {
	cc.deleteCalibracionHandler.Handle(ctx)
}
