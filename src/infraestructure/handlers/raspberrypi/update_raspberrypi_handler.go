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

func (h *UpdateRaspberrypiHandler) Handle(c *gin.Context) {
	var rpi entities.Raspberrypi

	// Obtener el ID de los parámetros
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Bind del JSON request
	if err := c.ShouldBindJSON(&rpi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Asignar el ID del path
	rpi.ID = id

	// Validar campos requeridos
	if rpi.Mac == "" || rpi.Modelo == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MAC y modelo son campos requeridos"})
		return
	}

	// Llamar al caso de uso
	err = h.raspberrypiUpdateUseCase.Execute(id, rpi)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rpi)
}
