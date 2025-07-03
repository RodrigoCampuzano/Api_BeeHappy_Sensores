package routes

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SensoresRoutes(r gin.IRouter, sensoresService *controllers.SensoresController) {
	// Grupo de sensores
	sensoresGroup := r.(*gin.RouterGroup).Group("/sensores")
	sensoresGroup.Use(middleware.AuthMiddleware())
	{
		// Operaciones CRUD básicas
		sensoresGroup.POST("/", sensoresService.CreateSensor)
		sensoresGroup.GET("/:id", sensoresService.GetSensor)
		sensoresGroup.PUT("/:id", sensoresService.UpdateSensor)
		sensoresGroup.DELETE("/:id", sensoresService.DeleteSensor)
		sensoresGroup.GET("/", sensoresService.GetAllSensores)
		sensoresGroup.GET("/raspberry/:id", sensoresService.GetByRaspberry)
		sensoresGroup.PUT("/:id/estado", sensoresService.UpdateEstadoSensores)
	}
}
