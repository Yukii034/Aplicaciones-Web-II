// =============================================================================
// SEMANA 3 - DÍA A - PUNTO DE PARTIDA DEL LIVE CODING
// =============================================================================
// Este es el estado del código al final de Semana 2 (variante con IDs).
// Todo vive en un único archivo: structs, slices, funciones y main.
//
// Durante la clase lo refactorizaremos a:
//   1. Paquetes separados (internal/inventario/)
//   2. Manejo de errores idiomático con (valor, error)
//   3. Interfaces como contrato entre paquetes
//
// Para ejecutar:   go run .
// =============================================================================

package main

import (
	"fmt"
	"semana3_vivo/internal/inventario"
) // Importa el modulo de inventario, guardado en la carpeta internal

// -----------------------------------------------------------------------------
// MAIN
// -----------------------------------------------------------------------------

func main() {
	// Agrega a la lista global - llama al modulo de inventario, con la funcion y struct especificos
	inventario.AgregarCategoria(inventario.Categoria{ID: 1, Nombre: "Bebidas"})
	inventario.AgregarCategoria(inventario.Categoria{ID: 2, Nombre: "Snacks"})

	var repo inventario.Repositorio = inventario.NewRepoMemoria() // Crea el nuevo repositorio. Cambiar si se usara una bd, se usan los mismos 3 metodos.

	// Guarda productos en el repositorio, llama a las funciones del repositorio y el struct especificos
	repo.Guardar(inventario.Producto{ID: 101, Nombre: "Agua 500ml", Precio: 0.50, Stock: 120, CategoriaID: 1})
	repo.Guardar(inventario.Producto{ID: 102, Nombre: "Cola 500ml", Precio: 1.00, Stock: 80, CategoriaID: 1})
	repo.Guardar(inventario.Producto{ID: 201, Nombre: "Papas fritas", Precio: 1.25, Stock: 45, CategoriaID: 2})

	// Usa las funciones del repositorio
	// Lista de productos
	fmt.Printf("Listando productos: \n")
	for _, p := range repo.Listar() {
		fmt.Printf("%+v \n", p)
	}

	// Busca un producto existente por ID
	p, err := repo.BuscarPorID(101)
	if err != nil {
		fmt.Printf("Error al buscar producto: %s \n", err.Error())
		return
	}
	fmt.Printf("\nEncontrado: %s\n", p.Nombre)

	// Busca uno que NO existe
	fantasma, err := repo.BuscarPorID(999)
	if err != nil {
		fmt.Printf("Error al buscar producto: %s \n", err.Error())
		return
	}
	fmt.Println(fantasma)
}
