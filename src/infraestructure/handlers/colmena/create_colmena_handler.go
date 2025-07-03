package handlers

import (
	"apisensores/src/application/colmena"
	"apisensores/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateColmenaHandler struct {
	colmenaCreateUseCase *colmena.CreateColmenaUseCase
}

func NewCreateColmenaHandler(colmenaService *colmena.CreateColmenaUseCase) *CreateColmenaHandler {
	return &CreateColmenaHandler{
		colmenaCreateUseCase: colmenaService,
	}
}

func (h *CreateColmenaHandler) Handle(ctx *gin.Context) {
	var colmenaData entities.Colmenas
	if err := ctx.ShouldBindJSON(&colmenaData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de colmena inválidos"})
		return
	}

	err := h.colmenaCreateUseCase.Execute(colmenaData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Colmena creada exitosamente"})
}
