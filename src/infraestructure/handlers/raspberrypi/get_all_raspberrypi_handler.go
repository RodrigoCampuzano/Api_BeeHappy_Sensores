package handlers

import (
	"apisensores/src/application/raspberrypi"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllRaspberrypiHandler struct {	
	raspberrypiGetAllUseCase *raspberrypi.GetAllRaspberryPiUseCase
}

func NewGetAllRaspberrypiHandler(raspberrypiService *raspberrypi.GetAllRaspberryPiUseCase) *GetAllRaspberrypiHandler {
	return &GetAllRaspberrypiHandler{
		raspberrypiGetAllUseCase: raspberrypiService,
	}
}

func (h *GetAllRaspberrypiHandler) Handle(c *gin.Context) {
	raspberrypis, err := h.raspberrypiGetAllUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	if len(raspberrypis) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron Raspberry Pi"})
		return
	}
	c.JSON(http.StatusOK, raspberrypis)
}