package routes

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func TipoSensoresRoutes(r gin.IRouter, tipoSensoresService *controllers.TipoSensoresController) {
	// Grupo de tipos de sensores
	tipoSensoresGroup := r.(*gin.RouterGroup).Group("/tipos-sensores")
	tipoSensoresGroup.Use(middleware.AuthMiddleware())
	{
		// Operaciones CRUD básicas
		tipoSensoresGroup.POST("/", tipoSensoresService.CreateTipoSensor)
		tipoSensoresGroup.GET("/:id", tipoSensoresService.GetTipoSensor)
		tipoSensoresGroup.PUT("/:id", tipoSensoresService.UpdateTipoSensor)
		tipoSensoresGroup.DELETE("/:id", tipoSensoresService.DeleteTipoSensor)
		tipoSensoresGroup.GET("/", tipoSensoresService.GetAllTipoSensores)               
		tipoSensoresGroup.GET("/nombre/:nombre", tipoSensoresService.GetByNombre) 
	}
}
