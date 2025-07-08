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
		raspberryGroup.POST("/", raspberryPiService.CreateRaspberryPi)
		raspberryGroup.GET("/:id", raspberryPiService.GetRaspberryPi)
		raspberryGroup.PUT("/:id", raspberryPiService.UpdateRaspberryPi)
		raspberryGroup.DELETE("/:id", raspberryPiService.DeleteRaspberryPi)
		raspberryGroup.GET("/", raspberryPiService.GetAllRaspberryPi)
		raspberryGroup.PUT("/:id/estado", raspberryPiService.UpdateEstadoRaspberryPi) 
		raspberryGroup.GET("/mac/:mac", raspberryPiService.GetByMAC)
	}
}
