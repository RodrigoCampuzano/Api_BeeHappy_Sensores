package handlers

import (
	"apisensores/src/application/colmena_raspberry"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateColmenaRaspberryHandler struct {
	colmenaRaspberryUpdateUseCase *colmena_raspberry.UpdateColmenaRaspberryUseCase
}

func NewUpdateColmenaRaspberryHandler(colmenaRaspberryService *colmena_raspberry.UpdateColmenaRaspberryUseCase) *UpdateColmenaRaspberryHandler {
	return &UpdateColmenaRaspberryHandler{
		colmenaRaspberryUpdateUseCase: colmenaRaspberryService,
	}
}

func (h *UpdateColmenaRaspberryHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var colmenaRaspberryData entities.ColmenaRaspberryPi
	if err := ctx.ShouldBindJSON(&colmenaRaspberryData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de colmena raspberry inválidos"})
		return
	}

	colmenaRaspberryData.ID = idInt
	err = h.colmenaRaspberryUpdateUseCase.Execute(colmenaRaspberryData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Colmena raspberry actualizada exitosamente"})
}
