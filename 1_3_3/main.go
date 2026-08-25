package main

import (
	"fmt"
)

//los punteros son basicamente los permisos que le das al programa para leer o editar directamente una variable
//se ocupa * cuando quieres leer solamente una variable
//se ocupa & cuando quieres editar y leer una variable

func main() {
	energia := 100
	var puntero *int = &energia //sacamos la direccion
	fmt.Println(energia)
	fmt.Println(*puntero) //hacemos print a lo que esta apuntando

	*puntero = 50        // cambiamos la direccion que esta apuntando
	fmt.Println(energia) // comprobamos que si cambio de verdad
}
