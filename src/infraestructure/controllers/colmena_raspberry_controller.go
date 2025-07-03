package controllers

import (
	"apisensores/src/application/colmena_raspberry"
	handlers "apisensores/src/infraestructure/handlers/colmena_raspberry"

	"github.com/gin-gonic/gin"
)

type ColmenaRaspberryController struct {
	createColmenaRaspberryHandler *handlers.CreateColmenaRaspberryHandler
	getColmenaRaspberryHandler    *handlers.GetColmenaRaspberryHandler
	updateColmenaRaspberryHandler *handlers.UpdateColmenaRaspberryHandler
	deleteColmenaRaspberryHandler *handlers.DeleteColmenaRaspberryHandler
}

func NewColmenaRaspberryController(
	createColmenaRaspberryUseCase *colmena_raspberry.CreateColmenaRaspberryUseCase,
	getColmenaRaspberryUseCase *colmena_raspberry.GetByColmenaUseCase,
	updateColmenaRaspberryUseCase *colmena_raspberry.UpdateColmenaRaspberryUseCase,
	deleteColmenaRaspberryUseCase *colmena_raspberry.DeleteColmenaRaspberryUseCase,
) *ColmenaRaspberryController {
	return &ColmenaRaspberryController{
		createColmenaRaspberryHandler: handlers.NewCreateColmenaRaspberryHandler(createColmenaRaspberryUseCase),
		getColmenaRaspberryHandler:    handlers.NewGetColmenaRaspberryHandler(getColmenaRaspberryUseCase),
		updateColmenaRaspberryHandler: handlers.NewUpdateColmenaRaspberryHandler(updateColmenaRaspberryUseCase),
		deleteColmenaRaspberryHandler: handlers.NewDeleteColmenaRaspberryHandler(deleteColmenaRaspberryUseCase),
	}
}

func (crc *ColmenaRaspberryController) CreateColmenaRaspberry(ctx *gin.Context) {
	crc.createColmenaRaspberryHandler.Handle(ctx)
}

func (crc *ColmenaRaspberryController) GetColmenaRaspberry(ctx *gin.Context) {
	crc.getColmenaRaspberryHandler.Handle(ctx)
}

func (crc *ColmenaRaspberryController) UpdateColmenaRaspberry(ctx *gin.Context) {
	crc.updateColmenaRaspberryHandler.Handle(ctx)
}

func (crc *ColmenaRaspberryController) DeleteColmenaRaspberry(ctx *gin.Context) {
	crc.deleteColmenaRaspberryHandler.Handle(ctx)
}