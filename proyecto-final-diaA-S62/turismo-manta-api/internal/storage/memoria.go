// Package storage gestiona el almacenamiento en memoria de los negocios.
package storage

import (
	"sync"

	"turismo-manta-api/internal/models"
)

// Memoria es un almacén de negocios que vive en memoria del proceso.
// Cada instancia tiene su propio slice y su propio mutex.
type Memoria struct {
	negocios []models.Negocio
	nextID   int
	mu       sync.Mutex
}

// NuevaMemoria crea un almacén vacío y listo para usar.
func NuevaMemoria() *Memoria {
	return &Memoria{
		negocios: []models.Negocio{},
		nextID:   1,
	}
}

// Seed carga los datos iniciales en memoria.
func (m *Memoria) Seed() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.negocios = []models.Negocio{
		{ID: 1, Nombre: "Hotel Oro Verde", Tipo: "hotel", Ciudad: "Manta"},
		{ID: 2, Nombre: "Restaurant El Marinero", Tipo: "restaurante", Ciudad: "Manta"},
		{ID: 3, Nombre: "Tour Operadora Pacífico", Tipo: "tour", Ciudad: "Manta"},
	}
	m.nextID = 4
}

// Listar devuelve todos los negocios.
func (m *Memoria) Listar() []models.Negocio {
	m.mu.Lock()
	defer m.mu.Unlock()

	copia := make([]models.Negocio, len(m.negocios))
	copy(copia, m.negocios)
	return copia
}

// BuscarPorID devuelve el negocio con el ID dado (patrón comma-ok).
func (m *Memoria) BuscarPorID(id int) (models.Negocio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, n := range m.negocios {
		if n.ID == id {
			return n, true
		}
	}
	return models.Negocio{}, false
}

// Crear agrega un negocio nuevo y devuelve el negocio con su ID asignado.
func (m *Memoria) Crear(n models.Negocio) models.Negocio {
	m.mu.Lock()
	defer m.mu.Unlock()

	n.ID = m.nextID
	m.nextID++
	m.negocios = append(m.negocios, n)
	return n
}

// Actualizar reemplaza el negocio con el ID dado.
func (m *Memoria) Actualizar(id int, datos models.Negocio) (models.Negocio, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.negocios {
		if n.ID == id {
			datos.ID = id
			m.negocios[i] = datos
			return datos, true
		}
	}
	return models.Negocio{}, false
}

// Borrar elimina el negocio con el ID dado.
func (m *Memoria) Borrar(id int) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.negocios {
		if n.ID == id {
			m.negocios = append(m.negocios[:i], m.negocios[i+1:]...)
			return true
		}
	}
	return false
}
