package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Informacion struct {
	Contrasena string `json:"password"`
	Usuario    string `json:"user"`
}

func main() {
	sv := gin.Default()

	sv.POST("/login", func(c *gin.Context) {
		var datos Informacion

		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON maluco"})
			return
		}
		fmt.Printf("el usuario se conecto con estos datos: ", datos.Usuario, datos.Contrasena)

		if datos.Usuario == "Benja" && datos.Contrasena == "123" {
			c.JSON(http.StatusOK, gin.H{"mensaje": "soy admin"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "pon bien la contraseña}"})
		}
	})

	sv.Run(":5050")
}
