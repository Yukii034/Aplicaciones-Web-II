// Package handlers contiene los handlers HTTP de la API.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"

	"turismo-manta-api/internal/models"
	"turismo-manta-api/internal/storage"
)

// Server agrupa las dependencias compartidas por los handlers.
// Cuando un handler necesita el almacenamiento, lo lee desde s.Storage
// en lugar de buscarlo en una variable global. Eso es Dependency Injection.
type Server struct {
	Storage *storage.Memoria
}

// NewServer construye un Server listo para usar.
// Quien llama (main) decide qué almacenamiento le inyecta.
func NewServer(s *storage.Memoria) *Server {
	return &Server{Storage: s}
}

// ListarNegocios atiende GET /api/v1/negocios.
func (s *Server) ListarNegocios(w http.ResponseWriter, _ *http.Request) {
	negocios := s.Storage.Listar()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(negocios); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ObtenerNegocio atiende GET /api/v1/negocios/{id}.
func (s *Server) ObtenerNegocio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	negocio, encontrado := s.Storage.BuscarPorID(id)
	if !encontrado {
		http.Error(w, "negocio no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(negocio); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// CrearNegocio atiende POST /api/v1/negocios.
func (s *Server) CrearNegocio(w http.ResponseWriter, r *http.Request) {
	var nuevo models.Negocio
	if err := json.NewDecoder(r.Body).Decode(&nuevo); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(nuevo.Nombre) == "" {
		http.Error(w, "el campo nombre es obligatorio", http.StatusBadRequest)
		return
	}

	creado := s.Storage.Crear(nuevo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(creado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// ActualizarNegocio atiende PUT /api/v1/negocios/{id}.
func (s *Server) ActualizarNegocio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	var datos models.Negocio
	if err := json.NewDecoder(r.Body).Decode(&datos); err != nil {
		http.Error(w, "JSON inválido: "+err.Error(), http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(datos.Nombre) == "" {
		http.Error(w, "el campo nombre es obligatorio", http.StatusBadRequest)
		return
	}

	actualizado, encontrado := s.Storage.Actualizar(id, datos)
	if !encontrado {
		http.Error(w, "negocio no encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(actualizado); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}

// BorrarNegocio atiende DELETE /api/v1/negocios/{id}.
func (s *Server) BorrarNegocio(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "id debe ser un número entero", http.StatusBadRequest)
		return
	}

	if !s.Storage.Borrar(id) {
		http.Error(w, "negocio no encontrado", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
