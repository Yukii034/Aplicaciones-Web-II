package main

import "fmt"

func main() {
	P1 := 49.99
	P2 := 29.99
	P3 := 15.50

	total := P1 + P2 + P3
	promedio := total / 3
	descuento := total * 0.85

	fmt.Printf("Producto 1: %.2f\n", P1)
	fmt.Printf("Producto 2: %.2f\n", P2)
	fmt.Printf("Producto 3: %.2f\n", P3)
	fmt.Printf("Total: %.2f\n", total)
	fmt.Printf("Promedio: %.2f\n", promedio)
	fmt.Printf("Descuento: %.2f\n", descuento)
}
