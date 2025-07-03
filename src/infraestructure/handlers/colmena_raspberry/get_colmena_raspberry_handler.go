package handlers

import (
	"apisensores/src/application/colmena_raspberry"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetColmenaRaspberryHandler struct {
	colmenaRaspberryGetUseCase *colmena_raspberry.GetByColmenaUseCase
}

func NewGetColmenaRaspberryHandler(colmenaRaspberryService *colmena_raspberry.GetByColmenaUseCase) *GetColmenaRaspberryHandler {
	return &GetColmenaRaspberryHandler{
		colmenaRaspberryGetUseCase: colmenaRaspberryService,
	}
}	

func (h *GetColmenaRaspberryHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	colmenaRaspberry, err := h.colmenaRaspberryGetUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if colmenaRaspberry.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Colmena Raspberry no encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, colmenaRaspberry)
}
