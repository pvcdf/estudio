package main

import (
	"1_3_11/pb"
	"fmt"
)

func main() {
	miMensaje := &pb.Humano{
		Nombre: "Benjamin Rojas",
		Edad:   21,
		Iq:     100,
	}

	fmt.Println(miMensaje)
}
