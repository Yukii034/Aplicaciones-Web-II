package service

import "errors"

var (
	ErrNombreVacio           = errors.New("Nombre es requerido")
	ErrPrecioNegativo        = errors.New("Precio no puede ser negativo")
	ErrNoEncontrado          = errors.New("Recurso no encontrado")
	ErrEmailEnUso            = errors.New("Email ya registrado")
	ErrCredencialesInvalidas = errors.New("Email o Contraseña incorrectos")
)
