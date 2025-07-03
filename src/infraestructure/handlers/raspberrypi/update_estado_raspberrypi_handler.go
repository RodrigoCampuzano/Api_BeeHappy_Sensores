package handlers

import (
	"apisensores/src/application/raspberrypi"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEstadoRaspberrypiHandler struct {	
	raspberrypiUpdateEstadoUseCase *raspberrypi.UpdateEstadoRaspberryPiUseCase
}

func NewUpdateEstadoRaspberrypiHandler(raspberrypiService *raspberrypi.UpdateEstadoRaspberryPiUseCase) *UpdateEstadoRaspberrypiHandler {
	return &UpdateEstadoRaspberrypiHandler{
		raspberrypiUpdateEstadoUseCase: raspberrypiService,
	}
}

func (h *UpdateEstadoRaspberrypiHandler) Handle(c *gin.Context) {	
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var raspberrypiData entities.Raspberrypi
	if err := c.ShouldBindJSON(&raspberrypiData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de Raspberry Pi inválidos"})
		return
	}

	raspberrypiData.ID = idInt
	err = h.raspberrypiUpdateEstadoUseCase.Execute(idInt, raspberrypiData.Estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado de Raspberry Pi actualizado exitosamente"})
}