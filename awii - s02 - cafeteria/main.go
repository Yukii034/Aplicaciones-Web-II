package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ── Helpers de lectura ──────────────────────────────────────────────────────

func leerLinea(lector *bufio.Reader) string {
	texto, _ := lector.ReadString('\n')
	return strings.TrimSpace(texto)
}

func leerEntero(lector *bufio.Reader, prompt string) int {
	for {
		fmt.Print(prompt)
		texto := leerLinea(lector)
		n, err := strconv.Atoi(texto)
		if err == nil {
			return n
		}
		fmt.Println("Ingresa un número entero válido.")
	}
}

func leerFloat(lector *bufio.Reader, prompt string) float64 {
	for {
		fmt.Print(prompt)
		texto := leerLinea(lector)
		f, err := strconv.ParseFloat(texto, 64)
		if err == nil {
			return f
		}
		fmt.Println("Ingresa un número decimal válido.")
	}
}

// ── Structs (sin cambios) ────────────────────────────────────────────────────

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

// ── Funciones existentes (sin cambios) ──────────────────────────────────────

func ListarClientes(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES ===")
	for _, c := range clientes {
		fmt.Printf("  ID: %d | %s | %s | Saldo: %.2f\n", c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func AgregarClientes(clientes []Cliente, cliente Cliente) []Cliente {
	return append(clientes, cliente)
}

func BuscarClientePorID(clientes []Cliente, id int) int {
	for i, c := range clientes {
		if c.ID == id {
			return i
		}
	}
	return -1
}

func ListarProductos(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS ===")
	for _, p := range productos {
		fmt.Printf("  ID: %d | %s | Precio: %.2f | Stock: %d | %s\n", p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}

func AgregarProducto(productos []Producto, producto Producto) []Producto {
	return append(productos, producto)
}

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, p := range productos {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func DescontarSaldo(cliente *Cliente, monto float64) error {
	if monto > cliente.Saldo {
		return fmt.Errorf("saldo insuficiente")
	}
	cliente.Saldo -= monto
	return nil
}

func DescontarStock(producto *Producto, cantidad int) error {
	if cantidad > producto.Stock {
		return fmt.Errorf("stock insuficiente")
	}
	producto.Stock -= cantidad
	return nil
}

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
	fecha string,
) ([]Pedido, error) {
	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, fmt.Errorf("cliente no encontrado")
	}
	idxP := BuscarProductoPorID(productos, productoID)
	if idxP == -1 {
		return pedidos, fmt.Errorf("producto no encontrado")
	}

	total := productos[idxP].Precio * float64(cantidad)

	if err := DescontarStock(&productos[idxP], cantidad); err != nil {
		return pedidos, err
	}
	if err := DescontarSaldo(&clientes[idxC], total); err != nil {
		return pedidos, err
	}

	nuevoPedido := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
		Fecha:      fecha,
	}
	return append(pedidos, nuevoPedido), nil
}

// ── Reporte cruzado (nuevo) ──────────────────────────────────────────────────

func PedidosDeCliente(pedidos []Pedido, clientes []Cliente, productos []Producto, clienteID int) {
	idxC := BuscarClientePorID(clientes, clienteID)
	if idxC == -1 {
		fmt.Println("Cliente no encontrado.")
		return
	}

	fmt.Printf("\n=== PEDIDOS DE: %s ===\n", clientes[idxC].Nombre)

	totalGastado := 0.0
	encontrado := false

	for _, p := range pedidos {
		if p.ClienteID != clienteID {
			continue
		}
		encontrado = true

		nombreProducto := "desconocido"
		idxP := BuscarProductoPorID(productos, p.ProductoID)
		if idxP != -1 {
			nombreProducto = productos[idxP].Nombre
		}

		fmt.Printf("  PedidoID: %d | Producto: %s | Cantidad: %d | Total: %.2f | Fecha: %s\n",
			p.ID, nombreProducto, p.Cantidad, p.Total, p.Fecha)
		totalGastado += p.Total
	}

	if !encontrado {
		fmt.Println("Este cliente no tiene pedidos.")
		return
	}

	fmt.Printf("  Total gastado: %.2f\n", totalGastado)
}

// ── main ─────────────────────────────────────────────────────────────────────

func main() {
	lector := bufio.NewReader(os.Stdin)

	clientes := []Cliente{
		{ID: 1, Nombre: "Cliente1", Carrera: "Carrera1", Saldo: 1000.00},
		{ID: 2, Nombre: "Cliente2", Carrera: "Carrera2", Saldo: 500.00},
		{ID: 3, Nombre: "Cliente3", Carrera: "Carrera3", Saldo: 750.00},
	}

	productos := []Producto{
		{ID: 1, Nombre: "Producto1", Precio: 10.00, Stock: 10, Categoria: "Categoria1"},
		{ID: 2, Nombre: "Producto2", Precio: 20.00, Stock: 20, Categoria: "Categoria2"},
		{ID: 3, Nombre: "Producto3", Precio: 30.00, Stock: 30, Categoria: "Categoria3"},
		{ID: 4, Nombre: "Producto4", Precio: 40.00, Stock: 40, Categoria: "Categoria4"},
	}

	pedidos := []Pedido{
		{ID: 1, ClienteID: 1, ProductoID: 1, Cantidad: 1, Total: 10.00, Fecha: "17-04-2026"},
		{ID: 2, ClienteID: 2, ProductoID: 2, Cantidad: 1, Total: 20.00, Fecha: "17-04-2026"},
	}

	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Listar clientes")
		fmt.Println("2. Listar productos")
		fmt.Println("3. Agregar cliente")
		fmt.Println("4. Agregar producto")
		fmt.Println("5. Registrar pedido")
		fmt.Println("6. Ver pedidos de un cliente")
		fmt.Println("0. Salir")

		opcion := leerEntero(lector, "Elige una opcion: ")

		switch opcion {

		case 1:
			ListarClientes(clientes)

		case 2:
			ListarProductos(productos)

		case 3:
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			saldo := leerFloat(lector, "Saldo inicial: ")
			clientes = AgregarClientes(clientes, Cliente{ID: id, Nombre: nombre, Carrera: carrera, Saldo: saldo})
			fmt.Println("Cliente agregado.")

		case 4:
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			precio := leerFloat(lector, "Precio: ")
			stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoria: ")
			categoria := leerLinea(lector)
			productos = AgregarProducto(productos, Producto{ID: id, Nombre: nombre, Precio: precio, Stock: stock, Categoria: categoria})
			fmt.Println("Producto agregado.")

		case 5:
			clienteID := leerEntero(lector, "ID del cliente: ")
			productoID := leerEntero(lector, "ID del producto: ")
			cantidad := leerEntero(lector, "Cantidad: ")
			fmt.Print("Fecha (dd-mm-aaaa): ")
			fecha := leerLinea(lector)
			var err error
			pedidos, err = RegistrarPedido(clientes, productos, pedidos, clienteID, productoID, cantidad, fecha)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Pedido registrado.")
			}

		case 6:
			clienteID := leerEntero(lector, "ID del cliente: ")
			PedidosDeCliente(pedidos, clientes, productos, clienteID)

		case 0:
			fmt.Println("Hasta luego!")
			return

		default:
			fmt.Println("Opcion no valida.")
		}
	}
}
