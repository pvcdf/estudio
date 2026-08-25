package main

import "fmt"

// tipos basicos, bool, strings etc

var booleano = false

var texto = "hola soy un string"

var enteros = int8(40) // si quieres numeros negativos, puedes poner int, pero nunca uint
//como dato curioso byte es lo mismo que int8 y rune a int32

var decimales = float32(3.14)

var complejos = complex(2, 3)

// tipos compuestos, arrays, slices, struct, etc

var arreglo [5]int // esto generara una lista de 5 numeros

var rebanada []int // esto generara una rebanada (slice), segun gemini este se usa mas que los arrays
//lo que entendi es que es un array que no esta obligado a un numero en especifico, puede ir creciendo a medida que le vayas poniendo datos

type Humano struct { // esto ta weno para hacer formularios (es como poo)
	Nombre string
	Edad   int
}

var mapa map[string]int // esto literalmente son los diccionarios en python, tienes una llave y un valor, en este caso la llave seria un string y el valor un numero entero

func main() {
	fmt.Println(booleano)
	fmt.Println(texto)
	fmt.Println(enteros)
	fmt.Println(complejos)
	fmt.Println(arreglo)
	fmt.Println(rebanada)
	rebanada = append(rebanada, 1)
	rebanada = append(rebanada, 2)
	rebanada = append(rebanada, 3)
	Estudiante := Humano{"Benjamin Rojas", 21}
	fmt.Println(rebanada)
	fmt.Println(Estudiante)
	fmt.Println(mapa)
}

// existe otro tipo de dato que son los punteros pero eso lo vamos a ver en investigar go: punteros :v
