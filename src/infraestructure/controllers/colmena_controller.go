package controllers

import (
	"apisensores/src/application/colmena"
	handlers "apisensores/src/infraestructure/handlers/colmena"

	"github.com/gin-gonic/gin"
)

type ColmenaController struct {
	createColmenaHandler *handlers.CreateColmenaHandler
	getColmenaHandler    *handlers.GetColmenaHandler
	updateColmenaHandler *handlers.UpdateColmenaHandler
	deleteColmenaHandler *handlers.DeleteColmenaHandler
	getAllColmenasHandler *handlers.GetAllColmenaHandler
	updateEstadoColmenaHandler *handlers.UpdateEstadoColmenaHandler
}

func NewColmenaController(
	createColmenaUseCase *colmena.CreateColmenaUseCase,
	getColmenaUseCase *colmena.GetColmenaUseCase,
	updateColmenaUseCase *colmena.UpdateColmenaUseCase,
	deleteColmenaUseCase *colmena.DeleteColmenaUseCase,
	getAllColmenasUseCase *colmena.GetAllColmenasUseCase,
	updateEstadoColmenaUseCase *colmena.UpdateEstadoColmenaUseCase,
) *ColmenaController {
	return &ColmenaController{
		createColmenaHandler: handlers.NewCreateColmenaHandler(createColmenaUseCase),
		getColmenaHandler:    handlers.NewGetColmenaHandler(getColmenaUseCase),
		updateColmenaHandler: handlers.NewUpdateColmenaHandler(updateColmenaUseCase),
		deleteColmenaHandler: handlers.NewDeleteColmenaHandler(deleteColmenaUseCase),
		getAllColmenasHandler: handlers.NewGetAllColmenaHandler(getAllColmenasUseCase),
		updateEstadoColmenaHandler: handlers.NewUpdateEstadoColmenaHandler(updateEstadoColmenaUseCase),
	}
}

func (cc *ColmenaController) CreateColmena(ctx *gin.Context) {
	cc.createColmenaHandler.Handle(ctx)
}

func (cc *ColmenaController) GetColmena(ctx *gin.Context) {
	cc.getColmenaHandler.Handle(ctx)
}

func (cc *ColmenaController) UpdateColmena(ctx *gin.Context) {
	cc.updateColmenaHandler.Handle(ctx)
}

func (cc *ColmenaController) DeleteColmena(ctx *gin.Context) {
	cc.deleteColmenaHandler.Handle(ctx)
}

func (cc *ColmenaController) GetAllColmenas(ctx *gin.Context) {
	cc.getAllColmenasHandler.Handle(ctx)
}

func (cc *ColmenaController) UpdateEstadoColmena(ctx *gin.Context) {
	cc.updateEstadoColmenaHandler.Handle(ctx)
}