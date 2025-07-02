package handlers

import (
	"apisensores/src/application/calibracion"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateCalibracionHandler struct {
	calibracionUpdateUseCase *calibracion.UpdateCalibracionUseCase
}

func NewUpdateCalibracionHandler(calibracionService *calibracion.UpdateCalibracionUseCase) *UpdateCalibracionHandler {
	return &UpdateCalibracionHandler{
		calibracionUpdateUseCase: calibracionService,
	}
}

func (h *UpdateCalibracionHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var calibracionData entities.Calibracion
	if err := ctx.ShouldBindJSON(&calibracionData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de calibración inválidos"})
		return
	}

	calibracionData.ID = idInt
	err = h.calibracionUpdateUseCase.Execute(calibracionData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Calibración actualizada exitosamente"})
}
