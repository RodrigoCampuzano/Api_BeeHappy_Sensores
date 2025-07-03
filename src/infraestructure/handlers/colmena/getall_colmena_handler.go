package handlers

import (
	"apisensores/src/application/colmena"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllColmenaHandler struct {
	colmenaService *colmena.GetAllColmenasUseCase
}

func NewGetAllColmenaHandler(colmenaService *colmena.GetAllColmenasUseCase) *GetAllColmenaHandler {		
	return &GetAllColmenaHandler{
		colmenaService: colmenaService,
	}
}

func (h *GetAllColmenaHandler) Handle(c *gin.Context) {
	colmenas, err := h.colmenaService.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})		
		return
	}
	c.JSON(http.StatusOK, colmenas)
}