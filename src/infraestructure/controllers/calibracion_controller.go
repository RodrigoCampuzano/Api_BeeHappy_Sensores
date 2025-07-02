package controllers

import (
	"apisensores/src/application/calibracion"
	handlerscalibracion "apisensores/src/infraestructure/handlers/calibracion"

	"github.com/gin-gonic/gin"
)

type CalibracionController struct {
	createCalibracionHandler *handlerscalibracion.CreateCalibracionHandler
	getCalibracionHandler    *handlerscalibracion.GetCalibracionHandler
	updateCalibracionHandler *handlerscalibracion.UpdateCalibracionHandler
	deleteCalibracionHandler *handlerscalibracion.DeleteCalibracionHandler	
}

func NewCalibracionController(
	createCalibracionUseCase *calibracion.CreateCalibracionUseCase,
	getCalibracionUseCase *calibracion.GetCalibracionUseCase,
	updateCalibracionUseCase *calibracion.UpdateCalibracionUseCase,
	deleteCalibracionUseCase *calibracion.DeleteCalibracionUseCase,
) *CalibracionController {
	return &CalibracionController{
		createCalibracionHandler: handlerscalibracion.NewCreateCalibracionHandler(createCalibracionUseCase),
		getCalibracionHandler:    handlerscalibracion.NewGetCalibracionHandler(getCalibracionUseCase),
		updateCalibracionHandler: handlerscalibracion.NewUpdateCalibracionHandler(updateCalibracionUseCase),
		deleteCalibracionHandler: handlerscalibracion.NewDeleteCalibracionHandler(deleteCalibracionUseCase),
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
