package inventario

import (
	"errors"
)

// -----------------------------------------------------------------------------
// ERRORES (importa errors)
// -----------------------------------------------------------------------------

var (
	ErrCategoriaNoEncontrada = errors.New("Categoria no encontrada")
	ErrProductoNoEncontrado  = errors.New("Producto no encontrado")
)

// -----------------------------------------------------------------------------
// INTERFAZ (conjunto de metodos)
// Analogía:: cualquier cosa que quiera llamarse Repositorio debe saber hacer estas tres operaciones
// Qué hace la interfaz en sí
// -----------------------------------------------------------------------------

type Repositorio interface {
	Guardar(p Producto) error
	BuscarPorID(id int) (Producto, error)
	Listar() []Producto
}

// -----------------------------------------------------------------------------
// STRUCTS (modelo con IDs — el tipo que luego verán en BD con GORM)
// -----------------------------------------------------------------------------

type RepoMemoria struct {
	productos []Producto
} // "BD" en memoria, se cambiaría por el repositorio de una BD real

type Categoria struct {
	ID     int
	Nombre string
}

type Producto struct {
	ID          int
	Nombre      string
	Precio      float64
	Stock       int
	CategoriaID int // referencia a Categoria por ID, NO anidación
}

// -----------------------------------------------------------------------------
// "BASE DE DATOS" EN MEMORIA
// -----------------------------------------------------------------------------

var categorias = []Categoria{}

// var productos = []Producto{}

// -----------------------------------------------------------------------------
// FUNCIONES DE REPOSITORIO
// -----------------------------------------------------------------------------

func NewRepoMemoria() *RepoMemoria {
	return &RepoMemoria{productos: []Producto{}}
} // Crea un Repositorio vacío, la & es el puntero para decir donde se encuentra

// -----------------------------------------------------------------------------
// FUNCIONES DE CATEGORÍAS
// -----------------------------------------------------------------------------

func AgregarCategoria(c Categoria) {
	categorias = append(categorias, c)
}

func BuscarCategoriaPorID(id int) (Categoria, error) {
	// PROBLEMA PEDAGÓGICO: si no existe, devolvemos Categoria{} (zero value).
	// No hay forma de que el llamador sepa si existía o no. En Semana 3
	// esto se arregla retornando (Categoria, error).
	for _, c := range categorias {
		if c.ID == id {
			return c, nil
		}
	}
	return Categoria{}, ErrCategoriaNoEncontrada
}

// -----------------------------------------------------------------------------
// FUNCIONES DE PRODUCTOS
// -----------------------------------------------------------------------------

func (r *RepoMemoria) Guardar(p Producto) error {
	r.productos = append(r.productos, p)
	return nil
} // La r indica que son funciones del repo

func (r *RepoMemoria) BuscarPorID(id int) (Producto, error) {
	// Mismo problema que BuscarCategoriaPorID.
	for _, p := range r.productos {
		if p.ID == id {
			return p, nil
		}
	}
	return Producto{}, ErrProductoNoEncontrado
}

func (r *RepoMemoria) Listar() []Producto {
	return r.productos
}
