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

func (rc *RaspberrypiController) CreateRaspberrypi(ctx *gin.Context) {
	rc.createRaspberrypiHandler.Handle(ctx)
}

func (rc *RaspberrypiController) GetRaspberrypi(ctx *gin.Context) {
	rc.getRaspberrypiHandler.Handle(ctx)
}

func (rc *RaspberrypiController) UpdateRaspberrypi(ctx *gin.Context) {
	rc.updateRaspberrypiHandler.Handle(ctx)
}

func (rc *RaspberrypiController) DeleteRaspberrypi(ctx *gin.Context) {
	rc.deleteRaspberrypiHandler.Handle(ctx)
}

func (rc *RaspberrypiController) GetAllRaspberrypi(ctx *gin.Context) {
	rc.getAllRaspberrypiHandler.Handle(ctx)
}

func (rc *RaspberrypiController) UpdateEstadoRaspberrypi(ctx *gin.Context) {
	rc.updateEstadoRaspberrypiHandler.Handle(ctx)
}

func (rc *RaspberrypiController) GetByMAC(ctx *gin.Context) {
	rc.getByMACHandler.Handle(ctx)
}