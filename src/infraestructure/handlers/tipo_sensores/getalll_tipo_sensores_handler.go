package handlers

import (
	"apisensores/src/application/tipo_sensores"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllTipoSensoresHandler struct {
	tipoSensoresRepository *tipo_sensores.GetAllTipoSensoresUseCase
}

func NewGetAllTipoSensoresHandler(tipoSensoresRepository *tipo_sensores.GetAllTipoSensoresUseCase) *GetAllTipoSensoresHandler {		
	return &GetAllTipoSensoresHandler{
		tipoSensoresRepository: tipoSensoresRepository,
	}
}

func (h *GetAllTipoSensoresHandler) Handle(c *gin.Context) {
	tipoSensores, err := h.tipoSensoresRepository.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tipoSensores)
}		