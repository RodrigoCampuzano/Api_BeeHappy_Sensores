package controllers

import (
	"apisensores/src/application/colmena"
	handlers "apisensores/src/infraestructure/handlers/colmena"

	"github.com/gin-gonic/gin"
)

type ColmenaController struct {
	createColmenaHandler       *handlers.CreateColmenaHandler
	getColmenaHandler          *handlers.GetColmenaHandler
	updateColmenaHandler       *handlers.UpdateColmenaHandler
	deleteColmenaHandler       *handlers.DeleteColmenaHandler
	getAllColmenasHandler      *handlers.GetAllColmenaHandler
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
		createColmenaHandler:       handlers.NewCreateColmenaHandler(createColmenaUseCase),
		getColmenaHandler:          handlers.NewGetColmenaHandler(getColmenaUseCase),
		updateColmenaHandler:       handlers.NewUpdateColmenaHandler(updateColmenaUseCase),
		deleteColmenaHandler:       handlers.NewDeleteColmenaHandler(deleteColmenaUseCase),
		getAllColmenasHandler:      handlers.NewGetAllColmenaHandler(getAllColmenasUseCase),
		updateEstadoColmenaHandler: handlers.NewUpdateEstadoColmenaHandler(updateEstadoColmenaUseCase),
	}
}
// @Summary Crear una nueva colmena
// @Description Crea un nuevo registro de colmena
// @Tags Colmena
// @Accept json
// @Produce json
// @Param request body entities.CreateColmenaRequest true "Datos de la colmena"
// @Success 201 {object} entities.ColmenaResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena [post]
func (cc *ColmenaController) CreateColmena(ctx *gin.Context) {
	cc.createColmenaHandler.Handle(ctx)
}

// @Summary Obtener una colmena
// @Description Obtiene los detalles de una colmena por su ID
// @Tags Colmena
// @Accept json
// @Produce json
// @Param id path int true "ID de la colmena"
// @Success 200 {object} entities.ColmenaResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena/{id} [get]
func (cc *ColmenaController) GetColmena(ctx *gin.Context) {
	cc.getColmenaHandler.Handle(ctx)
}

// @Summary Obtener todas las colmenas
// @Description Obtiene la lista de todas las colmenas registradas
// @Tags Colmena
// @Accept json
// @Produce json
// @Success 200 {object} entities.ColmenasListResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena [get]
func (cc *ColmenaController) GetAllColmenas(ctx *gin.Context) {
	cc.getAllColmenasHandler.Handle(ctx)
}

// @Summary Actualizar una colmena
// @Description Actualiza los datos de una colmena existente
// @Tags Colmena
// @Accept json
// @Produce json
// @Param id path int true "ID de la colmena"
// @Param request body entities.UpdateColmenaRequest true "Datos de actualización"
// @Success 200 {object} entities.ColmenaResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena/{id} [put]
func (cc *ColmenaController) UpdateColmena(ctx *gin.Context) {
	cc.updateColmenaHandler.Handle(ctx)
}

// @Summary Actualizar estado de colmena
// @Description Actualiza el estado de una colmena existente
// @Tags Colmena
// @Accept json
// @Produce json
// @Param id path int true "ID de la colmena"
// @Param request body entities.UpdateEstadoColmenaRequest true "Nuevo estado"
// @Success 200 {object} entities.ColmenaResponse
// @Failure 400 {object} entities.ErrorResponse
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena/{id}/estado [put]
func (cc *ColmenaController) UpdateEstadoColmena(ctx *gin.Context) {
	cc.updateEstadoColmenaHandler.Handle(ctx)
}

// @Summary Eliminar una colmena
// @Description Elimina una colmena por su ID
// @Tags Colmena
// @Accept json
// @Produce json
// @Param id path int true "ID de la colmena"
// @Success 204 "No Content"
// @Failure 404 {object} entities.ErrorResponse
// @Failure 500 {object} entities.ErrorResponse
// @Security ApiKeyAuth
// @Router /colmena/{id} [delete]
func (cc *ColmenaController) DeleteColmena(ctx *gin.Context) {
	cc.deleteColmenaHandler.Handle(ctx)
}
