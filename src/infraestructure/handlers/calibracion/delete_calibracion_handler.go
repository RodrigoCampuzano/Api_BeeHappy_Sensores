package handlers

import (
	"apisensores/src/application/calibracion"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteCalibracionHandler struct {
	calibracionDeleteUseCase *calibracion.DeleteCalibracionUseCase
}

func NewDeleteCalibracionHandler(calibracionService *calibracion.DeleteCalibracionUseCase) *DeleteCalibracionHandler {
	return &DeleteCalibracionHandler{
		calibracionDeleteUseCase: calibracionService,
	}
}

func (h *DeleteCalibracionHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.calibracionDeleteUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Calibración eliminada exitosamente"})
}
