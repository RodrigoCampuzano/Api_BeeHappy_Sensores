package handlers

import (
	"apisensores/src/application/colmena"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteColmenaHandler struct {
	colmenaDeleteUseCase *colmena.DeleteColmenaUseCase
}

func NewDeleteColmenaHandler(colmenaService *colmena.DeleteColmenaUseCase) *DeleteColmenaHandler {
	return &DeleteColmenaHandler{
		colmenaDeleteUseCase: colmenaService,
	}
}

func (h *DeleteColmenaHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.colmenaDeleteUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Colmena eliminada exitosamente"})
}
