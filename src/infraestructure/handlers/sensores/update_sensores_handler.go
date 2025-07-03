package handlers

import (
	"apisensores/src/application/sensores"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateSensoresHandler struct {
	sensoresUpdateUseCase *sensores.UpdateSensoresUseCase
}

func NewUpdateSensoresHandler(sensoresService *sensores.UpdateSensoresUseCase) *UpdateSensoresHandler {
	return &UpdateSensoresHandler{
		sensoresUpdateUseCase: sensoresService,
	}
}

func (h *UpdateSensoresHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var sensorData entities.Sensores
	if err := ctx.ShouldBindJSON(&sensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de sensor inválidos"})
		return
	}

	sensorData.ID = idInt
	err = h.sensoresUpdateUseCase.Execute(sensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sensor actualizado exitosamente"})
}
