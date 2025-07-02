package handlers

import (
	"apisensores/src/application/calibracion"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetCalibracionHandler struct {
	calibracionGetUseCase *calibracion.GetCalibracionUseCase
}

func NewGetCalibracionHandler(calibracionService *calibracion.GetCalibracionUseCase) *GetCalibracionHandler {
	return &GetCalibracionHandler{
		calibracionGetUseCase: calibracionService,
	}
}

func (h *GetCalibracionHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	calibracion, err := h.calibracionGetUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, calibracion)
}
