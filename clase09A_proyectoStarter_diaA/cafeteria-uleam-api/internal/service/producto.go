package service

import (
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)

type ProductoService struct {
	repo storage.ProductoRepository
}

func NewProductoService(repo storage.ProductoRepository) *ProductoService {
	return &ProductoService{repo: repo}
}

func (s *ProductoService) Listar() []models.Producto {
	return s.repo.ListarProductos()
}

func (s *ProductoService) Obtener(id int) (models.Producto, bool) {
	p, ok := s.repo.BuscarProductoPorID(id)
	if !ok {
		return models.Producto{}, false
	}
	return p, true
}

func (s *ProductoService) Crear(p models.Producto) (models.Producto, error) {
	if err := validarProducto(p); err != nil {
		return models.Producto{}, err
	}
	return s.repo.CrearProducto(p), nil
}

func (s *ProductoService) Actualizar(id int, p models.Producto) (models.Producto, error) {
	if err := validarProducto(p); err != nil {
		return models.Producto{}, err
	}
	p, ok := s.repo.ActualizarProducto(id, p)
	if !ok {
		return models.Producto{}, ErrNoEncontrado
	}
	return p, nil
}

func (s *ProductoService) Borrar(id int) error {
	if !s.repo.BorrarProducto(id) {
		return ErrNoEncontrado
	}
	return nil
}

func validarProducto(p models.Producto) error {
	if p.Nombre == "" {
		return ErrNombreVacio
	}
	if p.Precio < 0 {
		return ErrPrecioNegativo
	}
	return nil
}
