// Package models define las entidades del dominio turismo.
package models

// Negocio representa un negocio turístico de Manta.
type Negocio struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Tipo   string `json:"tipo"`
	Ciudad string `json:"ciudad"`
}
