package handlers

import (
	"apisensores/src/application/raspberrypi"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMacRaspberrypiHandler struct {	
	raspberrypiGetMacUseCase *raspberrypi.GetByMACUseCase	
}

func NewGetMacRaspberrypiHandler(raspberrypiService *raspberrypi.GetByMACUseCase) *GetMacRaspberrypiHandler {
	return &GetMacRaspberrypiHandler{
		raspberrypiGetMacUseCase: raspberrypiService,
	}
}

func (h *GetMacRaspberrypiHandler) Handle(c *gin.Context) {	
	mac := c.Param("mac")
	raspberrypi, err := h.raspberrypiGetMacUseCase.Execute(mac)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(raspberrypi) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Raspberry Pi no encontrada"})
		return
	}	

	if raspberrypi[0].ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Raspberry Pi no encontrada"})
		return
	}

	if raspberrypi[0].Mac == "" {	
		c.JSON(http.StatusNotFound, gin.H{"error": "MAC no encontrado"})
		return
	}

	c.JSON(http.StatusOK, raspberrypi)
}