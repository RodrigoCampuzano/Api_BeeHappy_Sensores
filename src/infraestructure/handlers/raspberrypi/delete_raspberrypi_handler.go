package handlers

import (
	"apisensores/src/application/raspberrypi"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteRaspberrypiHandler struct {
	raspberrypiDeleteUseCase *raspberrypi.DeleteRaspberrypiUseCase
}

func NewDeleteRaspberrypiHandler(raspberrypiService *raspberrypi.DeleteRaspberrypiUseCase) *DeleteRaspberrypiHandler {
	return &DeleteRaspberrypiHandler{
		raspberrypiDeleteUseCase: raspberrypiService,
	}
}

func (h *DeleteRaspberrypiHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.raspberrypiDeleteUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Raspberry pi eliminada exitosamente"})
}
