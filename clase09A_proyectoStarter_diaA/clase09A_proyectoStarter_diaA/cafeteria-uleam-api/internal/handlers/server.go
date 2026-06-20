package handlers

import (
	"cafeteria-uleam-api/internal/service"
)

type Server struct {
	Productos  *service.ProductoService
	Categorias *service.CategoriaService
	Auth       *service.AuthService
}

func NewServer(productos *service.ProductoService, categorias *service.CategoriaService, auth *service.AuthService) *Server {
	return &Server{
		Productos:  productos,
		Categorias: categorias,
		Auth:       auth,
	}
}
