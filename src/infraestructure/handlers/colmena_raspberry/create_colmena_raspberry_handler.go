package handlers

import (
	"apisensores/src/application/colmena_raspberry"
	"apisensores/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateColmenaRaspberryHandler struct {
	colmenaRaspberryCreateUseCase *colmena_raspberry.CreateColmenaRaspberryUseCase
}

func NewCreateColmenaRaspberryHandler(colmenaRaspberryService *colmena_raspberry.CreateColmenaRaspberryUseCase) *CreateColmenaRaspberryHandler {
	return &CreateColmenaRaspberryHandler{
		colmenaRaspberryCreateUseCase: colmenaRaspberryService,
	}
}

func (h *CreateColmenaRaspberryHandler) Handle(ctx *gin.Context) {
	var colmenaRaspberryData entities.ColmenaRaspberryPi
	if err := ctx.ShouldBindJSON(&colmenaRaspberryData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de colmena raspberry inválidos"})
		return
	}

	err := h.colmenaRaspberryCreateUseCase.Execute(colmenaRaspberryData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Colmena raspberry creada exitosamente"})
}
