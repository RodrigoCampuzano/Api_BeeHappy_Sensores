package handlers

import (
	"apisensores/src/application/sensores"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEstadoSensorHandler struct {	
	sensoresUpdateEstadoUseCase *sensores.UpdateEstadoSensorUseCase		
}

func NewUpdateEstadoSensorHandler(sensoresService *sensores.UpdateEstadoSensorUseCase) *UpdateEstadoSensorHandler {
	return &UpdateEstadoSensorHandler{
		sensoresUpdateEstadoUseCase: sensoresService,
	}
}

func (h *UpdateEstadoSensorHandler) Handle(c *gin.Context) {	
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var sensorData entities.Sensores
	if err := c.ShouldBindJSON(&sensorData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de sensor inválidos"})
		return
	}

	sensorData.ID = idInt
	err = h.sensoresUpdateEstadoUseCase.Execute(idInt, sensorData.Estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if sensorData.Estado == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estado inválido"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado de sensor actualizado exitosamente"})
}