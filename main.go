// Package main API de Sensores para Colmenas
package main

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Sensores para Colmenas
// @version 1.0
// @description API para la gestión de sensores en colmenas
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {
	r := gin.Default()

	// Middleware CORS
	r.Use(middleware.MiddlewareCORS())

	// Configuración de Swagger
	config := &ginSwagger.Config{
		URL:                      "http://localhost:8080/swagger/doc.json",
		DeepLinking:              true,
		DocExpansion:             "list",
		DefaultModelsExpandDepth: 1,
	}
	r.GET("/swagger/*any", ginSwagger.CustomWrapHandler(config, swaggerFiles.Handler))

	// Grupo base para la API v1
	api := r.Group("/api/v1")
	{
		// Inicializar rutas de cada módulo
		infraestructure.InitCalibracion(api)
		infraestructure.InitColmena(api)
		infraestructure.InitColmenaRaspberry(api)
		infraestructure.InitRaspberryPi(api)
		infraestructure.InitSensores(api)
		infraestructure.InitTipoSensores(api)
	}

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Println("Documentación Swagger disponible en http://localhost:8080/swagger/index.html")
	log.Fatal(r.Run(":8080"))
}
