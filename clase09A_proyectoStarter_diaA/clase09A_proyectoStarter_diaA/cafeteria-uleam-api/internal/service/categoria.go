package service

import (
	"cafeteria-uleam-api/internal/models"
	"cafeteria-uleam-api/internal/storage"
)

type CategoriaService struct {
	repo storage.CategoriaRepository
}

func NewCategoriaService(repo storage.CategoriaRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

func (s *CategoriaService) Listar() []models.Categoria {
	return s.repo.ListarCategorias()
}

func (s *CategoriaService) Obtener(id int) (models.Categoria, error) {
	c, ok := s.repo.BuscarCategoriaPorID(id)
	if !ok {
		return models.Categoria{}, ErrCategoriaNoEncontrada
	}
	return c, nil
}

func (s *CategoriaService) Crear(c models.Categoria) (models.Categoria, error) {
	if err := validarCategoria(c); err != nil {
		return models.Categoria{}, err
	}

	return s.repo.CrearCategoria(c), nil
}

func (s *CategoriaService) Actualizar(id int, c models.Categoria) (models.Categoria, error) {
	if err := validarCategoria(c); err != nil {
		return models.Categoria{}, err
	}

	c, ok := s.repo.ActualizarCategoria(id, c)
	if !ok {
		return models.Categoria{}, ErrCategoriaNoEncontrada
	}

	return c, nil
}

func (s *CategoriaService) Borrar(id int) error {
	if !s.repo.BorrarCategoria(id) {
		return ErrCategoriaNoEncontrada
	}

	return nil
}

func validarCategoria(c models.Categoria) error {
	if c.Nombre == "" {
		return ErrNombreVacio
	}

	return nil
}
