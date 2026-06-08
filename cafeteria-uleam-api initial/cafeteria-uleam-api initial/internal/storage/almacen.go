package storage

import "cafeteria-uleam-api/internal/models"

type Almacen interface{
	ListarProductos() []models.Producto
	BuscarProductoPorID(id int) (models.Producto, bool)
	CrearProducto(p models.Producto) models.Producto
	ActualizarProducto(id int, datos models.Producto) (models.Producto, bool)
	BorrarProducto(id int) bool

	ListarCategorias() []models.Categoria
	BuscarCategoriaPorID(id int) (models.Categoria, bool)
	CrearCategoria(c models.Categoria) models.Categoria
	ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool)
	BorrarCategoria(id int) bool
}

var _ Almacen = (*Memoria)(nil)