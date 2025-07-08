package routes

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func ColmenaRaspberryRoutes(r gin.IRouter, colmenaRaspberryService *controllers.ColmenaRaspberryController) {
	// Grupo de colmena-raspberry
	colmenaRaspberryGroup := r.(*gin.RouterGroup).Group("/colmena-raspberry")
	colmenaRaspberryGroup.Use(middleware.AuthMiddleware())
	{
		// Operaciones CRUD básicas
		colmenaRaspberryGroup.POST("/", colmenaRaspberryService.CreateColmenaRaspberry)
		colmenaRaspberryGroup.GET("/:id", colmenaRaspberryService.GetColmenaRaspberry)
		colmenaRaspberryGroup.PUT("/:id", colmenaRaspberryService.UpdateColmenaRaspberry)
		colmenaRaspberryGroup.DELETE("/:id", colmenaRaspberryService.DeleteColmenaRaspberry)
	}
}
