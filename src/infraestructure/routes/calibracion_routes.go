package routes

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func CalibracionRoutes(r gin.IRouter, calibracionService *controllers.CalibracionController) {
	// Grupo de calibración
	calibracionGroup := r.(*gin.RouterGroup).Group("/calibracion")
	calibracionGroup.Use(middleware.AuthMiddleware())
	{
		// Operaciones CRUD básicas
		calibracionGroup.POST("/", calibracionService.CreateCalibracion)
		calibracionGroup.GET("/:id", calibracionService.GetCalibracion)
		calibracionGroup.PUT("/:id", calibracionService.UpdateCalibracion)
		calibracionGroup.DELETE("/:id", calibracionService.DeleteCalibracion)
	}
}
