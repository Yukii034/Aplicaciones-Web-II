package main

import "fmt"

func main() {
	var nombre string
	var edad int

	fmt.Println("Escribe tu nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Escribe tu edad:")
	fmt.Scanln(&edad)

	fmt.Println(nombre, edad)
}
