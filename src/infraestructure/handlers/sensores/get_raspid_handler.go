package handlers

import (
	"apisensores/src/application/sensores"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type GetRaspidHandler struct {
	sensoresGetRaspidUseCase *sensores.GetByRaspberryUseCase
}

func NewGetRaspidHandler(sensoresService *sensores.GetByRaspberryUseCase) *GetRaspidHandler {
	return &GetRaspidHandler{
		sensoresGetRaspidUseCase: sensoresService,
	}
}

func (h *GetRaspidHandler) Handle(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	raspberrypi, err := h.sensoresGetRaspidUseCase.Execute(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(raspberrypi) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron sensores"})
		return
	}
	c.JSON(http.StatusOK, raspberrypi)
}