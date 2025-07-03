package controllers

import (
	"apisensores/src/application/calibracion"
	handlers "apisensores/src/infraestructure/handlers/calibracion"

	"github.com/gin-gonic/gin"
)

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

func (cc *CalibracionController) CreateCalibracion(ctx *gin.Context) {
	cc.createCalibracionHandler.Handle(ctx)
}

func (cc *CalibracionController) GetCalibracion(ctx *gin.Context) {
	cc.getCalibracionHandler.Handle(ctx)
}

func (cc *CalibracionController) UpdateCalibracion(ctx *gin.Context) {
	cc.updateCalibracionHandler.Handle(ctx)
}

func (cc *CalibracionController) DeleteCalibracion(ctx *gin.Context) {
	cc.deleteCalibracionHandler.Handle(ctx)
}
