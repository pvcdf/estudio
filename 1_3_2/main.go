package main

import "fmt"

type Humano struct {
	Nombre string
	Edad   int
	IQ     int
}

var Humanos []Humano //cree un slice para guardar a todos los humanos que pongamos

func crearHumano(nombre string, edad int, iq int) Humano { //ponemos aqui los argumentos que vamos a necesitar y que struct vamos a editar (Humano)
	return Humano{
		Nombre: nombre,
		Edad:   edad,
		IQ:     iq,
	}
}

func main() {
	for {
		var userName string
		var userAge int
		var userIQ int

		fmt.Println("pon tu nombre")
		fmt.Scanln(&userName)
		fmt.Println("pon tu edad")
		fmt.Scanln(&userAge)
		fmt.Println("pon tu iq")
		fmt.Scanln(&userIQ)

		Humanos = append(Humanos, crearHumano(userName, userAge, userIQ))
		fmt.Println(Humanos)
	}

}
