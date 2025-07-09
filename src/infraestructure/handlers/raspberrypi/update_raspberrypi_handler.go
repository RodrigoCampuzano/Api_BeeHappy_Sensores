package handlers

import (
	"apisensores/src/application/raspberrypi"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateRaspberrypiHandler struct {
	updateRaspberrypiUseCase *raspberrypi.UpdateRaspberryPiUseCase
}

func NewUpdateRaspberrypiHandler(useCase *raspberrypi.UpdateRaspberryPiUseCase) *UpdateRaspberrypiHandler {
	return &UpdateRaspberrypiHandler{
		updateRaspberrypiUseCase: useCase,
	}
}

func (h *UpdateRaspberrypiHandler) Handle(c *gin.Context) {
	// Obtener ID del path
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Obtener datos actuales del Raspberry Pi
	var rpi entities.Raspberrypi
	if err := c.ShouldBindJSON(&rpi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar el ID del path
	rpi.ID = id

	// Ejecutar el caso de uso
	err = h.updateRaspberrypiUseCase.Execute(id, rpi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rpi)
}