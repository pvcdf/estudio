package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Whitelist = "::1"

func guardiaUnimarc() gin.HandlerFunc { // aqui esta el middleware
	return func(c *gin.Context) {
		fmt.Println("se ta metiendo alguien")
		fmt.Println("su ip es", c.ClientIP())
		if c.ClientIP() != Whitelist {
			c.AbortWithStatusJSON(403, gin.H{"nope": "no eres admin viejo"})
			return
		}
		c.Next()
	}
}

func PermisoUsuario() gin.HandlerFunc { // aqui esta el middleware (version fruna)
	return func(c *gin.Context) {
		if c.ClientIP() == Whitelist {
			c.AbortWithStatusJSON(403, gin.H{"nope": "eres admin viejo"})
			return
		}
		c.Next()
	}
}

func main() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"exito": "pong!",
		})
	})
	server.GET("/admin", guardiaUnimarc(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hola": "bienvenido admin",
		})
	})
	server.GET("/usuario", PermisoUsuario(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hola": "bienvenido usuario (el admin no puede ver esto)",
		})
	})
	server.Run(":5050")
}
