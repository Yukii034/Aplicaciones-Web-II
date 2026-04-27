package main

import (
	"awii-s03-taller-relaciones/internal/cafeteria" //Importa el modulo
	"errors"
	"fmt"
)

func main() {

	var repo cafeteria.Repositorio = cafeteria.NewRepoMemoria()

	// Clientes añadidos
	repo.GuardarCliente(cafeteria.Cliente{ID: 1, Nombre: "Juan", Carrera: "TI", Saldo: 20})
	repo.GuardarCliente(cafeteria.Cliente{ID: 2, Nombre: "Ana", Carrera: "Civil", Saldo: 15})

	// Productos añadidos
	repo.GuardarProducto(cafeteria.Producto{ID: 1, Nombre: "Cola", Precio: 1.0, Stock: 10, Categoria: "bebidas"})
	repo.GuardarProducto(cafeteria.Producto{ID: 2, Nombre: "Jugo", Precio: 1.2, Stock: 8, Categoria: "bebidas"})
	repo.GuardarProducto(cafeteria.Producto{ID: 3, Nombre: "Pan", Precio: 0.5, Stock: 20, Categoria: "snacks"})

	//Obtener cliente existente
	fmt.Println("\n--- Buscar cliente existente ---")
	c, err := repo.ObtenerCliente(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Cliente encontrado:", c.Nombre)
	}

	//Cliente inexistente
	fmt.Println("\n--- Buscar cliente inexistente ---")
	_, err = repo.ObtenerCliente(99)
	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, cafeteria.ErrClienteNoEncontrado) {
			fmt.Println("Cliente no encontrado")
		}
	}

	//Listar productos
	fmt.Println("\n--- Lista de productos ---")
	for _, p := range repo.ListarProductos() {
		fmt.Printf("ID: %d - Producto: %s - Precio: $%.2f - Stock: %d\n",
			p.ID, p.Nombre, p.Precio, p.Stock)
	}

	//Mostrar Pedido
	fmt.Println("\n--- Pedido de cliente existente ---")

	cliente, _ := repo.ObtenerCliente(1)
	producto, _ := repo.ObtenerProducto(1)

	pedido := cafeteria.Pedido{
		ID:       1,
		Cliente:  cliente,
		Producto: producto,
		Total:    producto.Precio * 2,
		Fecha:    "2026-04-23",
	}

	fmt.Println("ID Pedido:", pedido.ID)
	fmt.Println("Cliente:", pedido.Cliente.Nombre)
	fmt.Println("Producto:", pedido.Producto.Nombre)
	fmt.Println("Total:", pedido.Total)
}
