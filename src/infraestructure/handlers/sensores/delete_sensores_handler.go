package handlers

import (
	"apisensores/src/application/sensores"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteSensoresHandler struct {
	sensoresDeleteUseCase *sensores.DeleteSensoresUseCase
}

func NewDeleteSensoresHandler(sensoresService *sensores.DeleteSensoresUseCase) *DeleteSensoresHandler {
	return &DeleteSensoresHandler{
		sensoresDeleteUseCase: sensoresService,
	}
}

func (h *DeleteSensoresHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.sensoresDeleteUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Sensor eliminado exitosamente"})
}
