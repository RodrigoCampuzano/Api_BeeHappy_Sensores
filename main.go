// Package main API de Sensores para Colmenas
package main

import (
	"apisensores/src/core/middleware"
	"apisensores/src/infraestructure"
	"log"

	"github.com/gin-gonic/gin"
)
func main() {
	r := gin.Default()

	// Middleware CORS
	r.Use(middleware.MiddlewareCORS())

	// Grupo base para la API v1
	api := r.Group("/api/v1")
	{
		// Inicializar rutas de cada módulo
		infraestructure.InitCalibracion(api)
	}

	log.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
