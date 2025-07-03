package handlers

import (
	"apisensores/src/application/tipo_sensores"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteTipoSensoresHandler struct {
	tipoSensoresDeleteUseCase *tipo_sensores.DeleteTipoSensoresUseCase
}

func NewDeleteTipoSensoresHandler(tipoSensoresService *tipo_sensores.DeleteTipoSensoresUseCase) *DeleteTipoSensoresHandler {
	return &DeleteTipoSensoresHandler{
		tipoSensoresDeleteUseCase: tipoSensoresService,
	}
}

func (h *DeleteTipoSensoresHandler) Handle(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.tipoSensoresDeleteUseCase.Execute(idInt)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Tipo de sensor eliminado exitosamente"})
}
