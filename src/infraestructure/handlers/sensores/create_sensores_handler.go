package handlers

import (
	"apisensores/src/application/sensores"
	"apisensores/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateSensoresHandler struct {
	sensoresCreateUseCase *sensores.CreateSensoresUseCase
}

func NewCreateSensoresHandler(sensoresService *sensores.CreateSensoresUseCase) *CreateSensoresHandler {
	return &CreateSensoresHandler{
		sensoresCreateUseCase: sensoresService,
	}
}

func (h *CreateSensoresHandler) Handle(ctx *gin.Context) {
	var sensorData entities.Sensores
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de sensor inválidos"})
		return
	}

	err := h.sensoresCreateUseCase.Execute(sensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Sensor creado exitosamente"})
}
