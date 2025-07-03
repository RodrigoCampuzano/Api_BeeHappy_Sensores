package handlers

import (
	"apisensores/src/application/tipo_sensores"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetTipoSensoresHandler struct {
	tipoSensoresGetUseCase *tipo_sensores.GetTipoSensoresUseCase
}

func NewGetTipoSensoresHandler(tipoSensoresService *tipo_sensores.GetTipoSensoresUseCase) *GetTipoSensoresHandler {
	return &GetTipoSensoresHandler{
		tipoSensoresGetUseCase: tipoSensoresService,
	}
}

func (h *GetTipoSensoresHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	tipoSensor, err := h.tipoSensoresGetUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tipoSensor)
}
