package handlers

import (
	"apisensores/src/application/tipo_sensores"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetByNombreTipoSensoresHandler struct {	
	tipoSensoresRepository *tipo_sensores.GetByNombreUseCase
}

func NewGetByNombreTipoSensoresHandler(tipoSensoresRepository *tipo_sensores.GetByNombreUseCase) *GetByNombreTipoSensoresHandler {
	return &GetByNombreTipoSensoresHandler{
		tipoSensoresRepository: tipoSensoresRepository,
	}

}

func (h *GetByNombreTipoSensoresHandler) Handle(c *gin.Context) {	
	nombre := c.Param("nombre")
	tipoSensores, err := h.tipoSensoresRepository.Execute(nombre)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	if tipoSensores.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tipo de sensor no encontrado"})
		return
	}
	if tipoSensores.Nombre == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tipo de sensor no encontrado"})
		return
	}
	c.JSON(http.StatusOK, tipoSensores)
}