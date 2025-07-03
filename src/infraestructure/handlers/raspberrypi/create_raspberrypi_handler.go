package handlers

import (
	"apisensores/src/application/raspberrypi"
	"apisensores/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateRaspberrypiHandler struct {
	raspberrypiCreateUseCase *raspberrypi.CreateRaspberrypiUseCase
}

func NewCreateRaspberrypiHandler(raspberrypiService *raspberrypi.CreateRaspberrypiUseCase) *CreateRaspberrypiHandler {
	return &CreateRaspberrypiHandler{
		raspberrypiCreateUseCase: raspberrypiService,
	}
}

func (h *CreateRaspberrypiHandler) Handle(ctx *gin.Context) {
	var raspberrypiData entities.Raspberrypi
	if err := ctx.ShouldBindJSON(&raspberrypiData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos de raspberry pi inválidos"})
		return
	}

	err := h.raspberrypiCreateUseCase.Execute(raspberrypiData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Raspberry pi creada exitosamente"})
} 