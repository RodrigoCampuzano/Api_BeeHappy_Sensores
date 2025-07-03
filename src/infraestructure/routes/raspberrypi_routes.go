package routes

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func RaspberryPiRoutes(r gin.IRouter, raspberryPiService *controllers.RaspberrypiController) {
	// Grupo de raspberry pi
	raspberryGroup := r.(*gin.RouterGroup).Group("/raspberry")
	raspberryGroup.Use(middleware.AuthMiddleware())
	{
		// Operaciones CRUD básicas 	
		raspberryGroup.POST("/", raspberryPiService.CreateRaspberrypi)
		raspberryGroup.GET("/:id", raspberryPiService.GetRaspberrypi)
		raspberryGroup.PUT("/:id", raspberryPiService.UpdateRaspberrypi)
		raspberryGroup.DELETE("/:id", raspberryPiService.DeleteRaspberrypi)
		raspberryGroup.GET("/", raspberryPiService.GetAllRaspberrypi)
		raspberryGroup.PUT("/:id/estado", raspberryPiService.UpdateEstadoRaspberrypi) 
		raspberryGroup.GET("/mac/:mac", raspberryPiService.GetByMAC)
	}
}
