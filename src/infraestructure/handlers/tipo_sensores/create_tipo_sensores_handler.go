package handlers

import (
	"apisensores/src/application/tipo_sensores"
	"apisensores/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTipoSensoresHandler struct {
	tipoSensoresCreateUseCase *tipo_sensores.CreateTipoSensoresUseCase
}

func NewCreateTipoSensoresHandler(tipoSensoresService *tipo_sensores.CreateTipoSensoresUseCase) *CreateTipoSensoresHandler {
	return &CreateTipoSensoresHandler{
		tipoSensoresCreateUseCase: tipoSensoresService,
	}
}

func (h *CreateTipoSensoresHandler) Handle(ctx *gin.Context) {
	var tipoSensorData entities.Tipo_Sesnores
	if err := ctx.ShouldBindJSON(&tipoSensorData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de tipo de sensor inválidos"})
		return
	}

	err := h.tipoSensoresCreateUseCase.Execute(tipoSensorData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Tipo de sensor creado exitosamente"})
}
