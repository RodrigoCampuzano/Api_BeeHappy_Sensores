package routes

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func ColmenaRoutes(r gin.IRouter, colmenaService *controllers.ColmenaController) {
	// Grupo de colmenas
	colmenaGroup := r.(*gin.RouterGroup).Group("/colmenas")
	colmenaGroup.Use(middleware.AuthMiddleware())
	{
		// Operaciones CRUD básicas
		colmenaGroup.POST("/", colmenaService.CreateColmena)
		colmenaGroup.GET("/:id", colmenaService.GetColmena)
		colmenaGroup.PUT("/:id", colmenaService.UpdateColmena)
		colmenaGroup.DELETE("/:id", colmenaService.DeleteColmena)
		colmenaGroup.GET("/", colmenaService.GetColmena)
		colmenaGroup.PUT("/:id/estado", colmenaService.UpdateColmena) 
	}
}
