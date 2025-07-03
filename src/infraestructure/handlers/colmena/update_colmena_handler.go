package handlers

import (
	"apisensores/src/application/colmena"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateColmenaHandler struct {
	colmenaUpdateUseCase *colmena.UpdateColmenaUseCase
}

func NewUpdateColmenaHandler(colmenaService *colmena.UpdateColmenaUseCase) *UpdateColmenaHandler {
	return &UpdateColmenaHandler{
		colmenaUpdateUseCase: colmenaService,
	}
}

func (h *UpdateColmenaHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var colmenaData entities.Colmenas
	if err := ctx.ShouldBindJSON(&colmenaData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de colmena inválidos"})
		return
	}

	colmenaData.ID = idInt
	err = h.colmenaUpdateUseCase.Execute(colmenaData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Colmena actualizada exitosamente"})
}
