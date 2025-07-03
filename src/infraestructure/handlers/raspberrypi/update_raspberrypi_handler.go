package handlers

import (
	"apisensores/src/application/raspberrypi"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRaspberrypiHandler struct {
	raspberrypiUpdateUseCase *raspberrypi.UpdateRaspberryPiUseCase
}

func NewUpdateRaspberrypiHandler(raspberrypiService *raspberrypi.UpdateRaspberryPiUseCase) *UpdateRaspberrypiHandler {
	return &UpdateRaspberrypiHandler{
		raspberrypiUpdateUseCase: raspberrypiService,
	}
}

func (h *UpdateRaspberrypiHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var raspberrypiData entities.Raspberrypi
	if err := ctx.ShouldBindJSON(&raspberrypiData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de raspberry pi inválidos"})
		return
	}

	raspberrypiData.ID = idInt
	err = h.raspberrypiUpdateUseCase.Execute(raspberrypiData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Raspberry pi actualizada exitosamente"})
}
