package handlers

import (
	"apisensores/src/application/colmena"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetColmenaHandler struct {
	colmenaGetUseCase *colmena.GetColmenaUseCase
}

func NewGetColmenaHandler(colmenaService *colmena.GetColmenaUseCase) *GetColmenaHandler {
	return &GetColmenaHandler{
		colmenaGetUseCase: colmenaService,
	}
}

func (h *GetColmenaHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	colmena, err := h.colmenaGetUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, colmena)
}
