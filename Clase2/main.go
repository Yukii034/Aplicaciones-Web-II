package main

import "fmt"

type Producto struct { // como objeto con atributos (funciona como array)
	Nombre string
	Precio float64
	Stock  int64
}

func agregarProducto(productos []Producto, producto Producto) []Producto {
	return append(productos, producto)
}

func calcularTotal(productos []Producto) float64 { //se nos envia el struct y la funcion devuelve el total en float64
	total := 0.0
	for _, producto := range productos { //recorre el struct (como un array estático), con el tamaño actual del mismo, no tiene un indice
		total += producto.Precio * float64(producto.Stock) //convierte el int64 a float64 del stock y lo multiplica con el precio
	}
	return total
}

func buscarProducto(productos []Producto, nombre string) (Producto, bool) { //se nos envia el struct y el nombre tipo string, la funcion devuelve el producto (como array) y el tipo de dato bool
	for _, producto := range productos {
		if producto.Nombre == nombre {
			return producto, true
		}
	}
	return Producto{}, false
}

func main() {
	productos := []Producto{ //productos estáticos del struct
		{Nombre: "Producto1", Precio: 10.00, Stock: 10},
		{Nombre: "Producto2", Precio: 20.00, Stock: 20},
	}
	producto := Producto{Nombre: "Producto3", Precio: 30.00, Stock: 30}
	productos = agregarProducto(productos, producto) //agg un producto al struct
	total := calcularTotal(productos)
	fmt.Println("Total:", total)

	producto, ok := buscarProducto(productos, "Producto3") //ok es como un bool
	if ok {
		fmt.Println("Producto encontrado", producto)
	} else {
		fmt.Println("Producto no encontrado")
	}
}
