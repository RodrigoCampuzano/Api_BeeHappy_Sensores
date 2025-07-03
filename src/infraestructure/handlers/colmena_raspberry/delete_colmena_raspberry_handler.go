package handlers

import (
	"apisensores/src/application/colmena_raspberry"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteColmenaRaspberryHandler struct {
	colmenaRaspberryDeleteUseCase *colmena_raspberry.DeleteColmenaRaspberryUseCase
}

func NewDeleteColmenaRaspberryHandler(colmenaRaspberryService *colmena_raspberry.DeleteColmenaRaspberryUseCase) *DeleteColmenaRaspberryHandler {
	return &DeleteColmenaRaspberryHandler{
		colmenaRaspberryDeleteUseCase: colmenaRaspberryService,
	}
}

func (h *DeleteColmenaRaspberryHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.colmenaRaspberryDeleteUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Colmena raspberry eliminada exitosamente"})
}
