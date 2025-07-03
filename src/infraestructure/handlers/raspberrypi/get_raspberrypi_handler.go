package handlers

import (
	"apisensores/src/application/raspberrypi"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetRaspberrypiHandler struct {
	raspberrypiGetUseCase *raspberrypi.GetRaspberryPiUseCase
}

func NewGetRaspberrypiHandler(raspberrypiService *raspberrypi.GetRaspberryPiUseCase) *GetRaspberrypiHandler {
	return &GetRaspberrypiHandler{
		raspberrypiGetUseCase: raspberrypiService,
	}
}

func (h *GetRaspberrypiHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	raspberrypi, err := h.raspberrypiGetUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, raspberrypi)
}
