package main

import "fmt"

func main() {
	nombre := "Pierina"
	edad := 20
	carrera := "TI"
	semestre := 6
	promedio := 9.00

	fmt.Printf("Mi nombre es: %s, tengo %d años\n", nombre, edad)
	fmt.Printf("Estudio %s, semestre %d\n", carrera, semestre)
	fmt.Printf("Mi promedio es: %.2f\n", promedio)
}
