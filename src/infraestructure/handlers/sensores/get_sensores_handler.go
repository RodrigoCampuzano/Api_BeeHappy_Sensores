package handlers

import (
	"apisensores/src/application/sensores"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetSensoresHandler struct {
	sensoresGetUseCase *sensores.GetSensoresUseCase
}

func NewGetSensoresHandler(sensoresService *sensores.GetSensoresUseCase) *GetSensoresHandler {
	return &GetSensoresHandler{
		sensoresGetUseCase: sensoresService,
	}
}

func (h *GetSensoresHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	sensor, err := h.sensoresGetUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sensor)
}
