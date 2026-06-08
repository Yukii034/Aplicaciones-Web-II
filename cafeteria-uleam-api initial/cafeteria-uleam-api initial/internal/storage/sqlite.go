package storage

import (
	"cafeteria-uleam-api/internal/models"

	"gorm.io/gorm"
)

type AlmacenSQLite struct {
	db *gorm.DB
}

func NewAlmacenSQLite(db *gorm.DB) *AlmacenSQLite {
	return &AlmacenSQLite{db: db}
}

func (a *AlmacenSQLite) ListarProductos() []models.Producto {
	var productos []models.Producto
	a.db.Find(&productos) // select * from productos
	return productos
}

func (a *AlmacenSQLite) BuscarProductoPorID(id int) (models.Producto, bool) {
	var producto models.Producto
	if err := a.db.First(&producto, id).Error; err != nil {
		return models.Producto{}, false
	}
	return producto, true
}

func (a *AlmacenSQLite) CrearProducto(p models.Producto) models.Producto {
	a.db.Create(&p) // insert into productos
	return p
}

func (a *AlmacenSQLite) ActualizarProducto(id int, datos models.Producto) (models.Producto, bool) {
	var productoExistente models.Producto
	if err := a.db.First(&productoExistente, id).Error; err != nil {
		return models.Producto{}, false
	}
	datos.ID = id
	a.db.Save(&datos) // update productos set nombre = ?, precio = ?, stock = ?, categoriaID = ?
	return datos, true
}

func (a *AlmacenSQLite) BorrarProducto(id int) bool {
	res := a.db.Delete(&models.Producto{}, id) // delete from productos where id =
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) ListarCategorias() []models.Categoria {
	var categorias []models.Categoria
	a.db.Find(&categorias) // select * from categorias
	return categorias
}

func (a *AlmacenSQLite) BuscarCategoriaPorID(id int) (models.Categoria, bool) {
	var categoria models.Categoria
	if err := a.db.First(&categoria, id).Error; err != nil {
		return models.Categoria{}, false
	}
	return categoria, true
}

func (a *AlmacenSQLite) CrearCategoria(c models.Categoria) models.Categoria {
	a.db.Create(&c) // insert into categorias
	return c
}

func (a *AlmacenSQLite) ActualizarCategoria(id int, datos models.Categoria) (models.Categoria, bool) {
	var categoriaExistente models.Producto
	if err := a.db.First(&categoriaExistente, id).Error; err != nil {
		return models.Categoria{}, false
	}
	datos.ID = id
	a.db.Save(&datos) // update categorias set nombre = ?, precio = ?, stock = ?, categoriaID = ?
	return datos, true
}

func (a *AlmacenSQLite) BorrarCategoria(id int) bool {
	res := a.db.Delete(&models.Categoria{}, id) // delete from categorias where id =
	return res.RowsAffected > 0
}

func (a *AlmacenSQLite) SembrarVacio() {
	var n int64
	a.db.Model(&models.Categoria{}).Count(&n)
	if n > 0 {
		return
	}
	categorias := []models.Categoria{
		{Nombre: "Categoria 1", Descripcion: "Descripcion Categoria 1"},
		{Nombre: "Categoria 2", Descripcion: "Descripcion Categoria 2"},
		{Nombre: "Categoria 3", Descripcion: "Descripcion Categoria 3"},
	}
	a.db.Create(&categorias)
	productos := []models.Producto{
		{Nombre: "Producto 1", Precio: 10.0, Stock: 2, CategoriaID: 1},
		{Nombre: "Producto 2", Precio: 20.0, Stock: 4, CategoriaID: 2},
		{Nombre: "Producto 3", Precio: 30.0, Stock: 6, CategoriaID: 3},
	}
	a.db.Create(&productos)
}
