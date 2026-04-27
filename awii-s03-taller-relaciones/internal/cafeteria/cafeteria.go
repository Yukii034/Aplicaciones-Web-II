package cafeteria

import (
	"errors"
)

// -----------------------------------------------------------------------------
// ERRORES (importa errors)
// -----------------------------------------------------------------------------

var (
	ErrClienteNoEncontrado  = errors.New("Cliente no encontrado")
	ErrProductoNoEncontrado = errors.New("Producto no encontrado")
	ErrStockInsuficiente    = errors.New("Stock insuficiente")
	ErrSaldoInsuficiente    = errors.New("Saldo insuficiente del cliente")
)

// -----------------------------------------------------------------------------
// INTERFAZ (conjunto de metodos)
// Analogía:: cualquier cosa que quiera llamarse Repositorio debe saber hacer estas tres operaciones
// Qué hace la interfaz en sí
// -----------------------------------------------------------------------------

type Repositorio interface {
	GuardarCliente(cliente Cliente) error
	ObtenerCliente(id int) (Cliente, error)
	ListarClientes() []Cliente
	GuardarProducto(producto Producto) error
	ObtenerProducto(id int) (Producto, error)
	ListarProductos() []Producto
}

// ── Structs ────────────────────────────────────────────────────

type RepoMemoria struct {
	clientes  []Cliente
	productos []Producto
	pedidos   []Pedido
} // "BD" en memoria, se cambiaría por el repositorio de una BD real

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
	ID       int
	Cliente  Cliente
	Producto Producto
	Total    float64
	Fecha    string
}

// ── Funciones ──────────────────────────────────────

// Clientes
func (r *RepoMemoria) GuardarCliente(c Cliente) error {
	r.clientes = append(r.clientes, c)
	return nil
}

func (r *RepoMemoria) ObtenerCliente(id int) (Cliente, error) {
	for _, c := range r.clientes {
		if c.ID == id {
			return c, nil
		}
	}
	return Cliente{}, ErrClienteNoEncontrado
}

func (r *RepoMemoria) ListarClientes() []Cliente {
	return r.clientes
}

// Productos
func (r *RepoMemoria) ListarProductos() []Producto {
	return r.productos
}

func (r *RepoMemoria) ObtenerProducto(id int) (Producto, error) {
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

func (r *RepoMemoria) GuardarProducto(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
}

// -----------------------------------------------------------------------------
// FUNCIONES DE REPOSITORIO
// -----------------------------------------------------------------------------

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{}
} // Crea un Repositorio vacío, la & es el puntero para decir donde se encuentra

// Verificación en tiempo de compilación:
// Si RepoMemoria NO cumple Repository, esto da error al compilar.
var _ Repositorio = (*RepoMemoria)(nil)
