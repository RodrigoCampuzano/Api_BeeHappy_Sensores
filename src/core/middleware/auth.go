package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, gin.H{"error": "No se proporcionó token de autorización"})
			c.Abort()
			return
		}

		// Verificar formato "Bearer {token}"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(401, gin.H{"error": "Formato de token inválido"})
			c.Abort()
			return
		}

		// Aquí puedes agregar la lógica para validar el token con tu API de autenticación
		// Por ejemplo, hacer una petición al servicio en el puerto 8081

		c.Next()
	}
}
