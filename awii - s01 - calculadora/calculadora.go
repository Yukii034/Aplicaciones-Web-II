package main

import "fmt"

func main() {
	var a float64
	var b float64
	var op string

	fmt.Println("Ingresa el primer número:")
	fmt.Scanln(&a)
	fmt.Println("Ingresa el segundo número:")
	fmt.Scanln(&b)
	fmt.Println("Ingresa la operación (+, -, *, /, ^, !):")
	fmt.Scanln(&op)

	fmt.Println("=====CALCULADORA CIENTÍFICA v1.0=====")
	switch op {
	case "+":
		suma := a + b
		fmt.Printf("El resultado de la suma es: %.2f\n", suma)

	case "-":
		resta := a - b
		fmt.Printf("El resultado de la resta es: %.2f\n", resta)

	case "*":
		multi := a * b
		fmt.Printf("El resultado de la multiplicación es: %.2f\n", multi)

	case "/":
		divi := a / b
		fmt.Printf("El resultado de la división es: %.2f\n", divi)

	case "^":
		pot := 1.0
		i := 0.00
		for {
			if i == b {
				break
			}
			pot *= a
			i++
		}
		fmt.Printf("El resultado de la potencia es: %.2f\n", pot)

	case "!":
		fac := 1.00
		for i := range int64(a) {
			fac *= float64(i) // Usar float64 para convertir el for en float
		}
		fmt.Printf("El resultado del factorial es: %.2f\n", fac)

	default:
		fmt.Println("Operación inválida")
	}
}
