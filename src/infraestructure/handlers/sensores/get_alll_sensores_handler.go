package handlers

import (
	"apisensores/src/application/sensores"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllSensoresHandler struct {
	sensoresGetAllUseCase *sensores.GetAllSensoresUseCase
}

func NewGetAllSensoresHandler(sensoresService *sensores.GetAllSensoresUseCase) *GetAllSensoresHandler {
	return &GetAllSensoresHandler{
		sensoresGetAllUseCase: sensoresService,
	}
}

func (h *GetAllSensoresHandler) Handle(c *gin.Context) {	
	sensores, err := h.sensoresGetAllUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(sensores) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No se encontraron sensores"})
		return
	}
	c.JSON(http.StatusOK, sensores)
}