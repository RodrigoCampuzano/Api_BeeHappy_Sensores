package handlers

import (
	"apisensores/src/application/colmena"
	"apisensores/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateEstadoColmenaHandler struct {
	colmenaUpdateEstadoUseCase *colmena.UpdateEstadoColmenaUseCase
}

func NewUpdateEstadoColmenaHandler(colmenaService *colmena.UpdateEstadoColmenaUseCase) *UpdateEstadoColmenaHandler {
	return &UpdateEstadoColmenaHandler{
		colmenaUpdateEstadoUseCase: colmenaService,
	}
}

func (h *UpdateEstadoColmenaHandler) Handle(c *gin.Context) {	
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var colmenaData entities.Colmenas
	if err := c.ShouldBindJSON(&colmenaData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de colmena inválidos"})
		return
	}

	colmenaData.ID = idInt
	err = h.colmenaUpdateEstadoUseCase.Execute(idInt, colmenaData.Estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Estado de colmena actualizado exitosamente"})
	
}