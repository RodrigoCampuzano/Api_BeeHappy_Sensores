package handlers

import (
	"apisensores/src/application/tipo_sensores"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateTipoSensoresHandler struct {
	tipoSensoresUpdateUseCase *tipo_sensores.UpdateTipoSensoresUseCase
}

func NewUpdateTipoSensoresHandler(tipoSensoresService *tipo_sensores.UpdateTipoSensoresUseCase) *UpdateTipoSensoresHandler {
	return &UpdateTipoSensoresHandler{
		tipoSensoresUpdateUseCase: tipoSensoresService,
	}
}

func (h *UpdateTipoSensoresHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var tipoSensorData entities.Tipo_Sesnores
	if err := ctx.ShouldBindJSON(&tipoSensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de tipo de sensor inválidos"})
		return
	}

	tipoSensorData.ID = idInt
	err = h.tipoSensoresUpdateUseCase.Execute(tipoSensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tipo de sensor actualizado exitosamente"})
}
