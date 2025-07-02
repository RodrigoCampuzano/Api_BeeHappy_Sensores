package handlers

import (
	"apisensores/src/application/calibracion"
	"apisensores/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateCalibracionHandler struct {
	calibracionCreateUseCase *calibracion.CreateCalibracionUseCase
}

func NewCreateCalibracionHandler(calibracionService *calibracion.CreateCalibracionUseCase) *CreateCalibracionHandler {
	return &CreateCalibracionHandler{
		calibracionCreateUseCase: calibracionService,
	}
}

func (h *CreateCalibracionHandler) Handle(ctx *gin.Context) {
	var calibracionData entities.Calibracion
	if err := ctx.ShouldBindJSON(&calibracionData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de calibración inválidos"})
		return
	}

	err := h.calibracionCreateUseCase.Execute(calibracionData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Calibración creada exitosamente"})
}
