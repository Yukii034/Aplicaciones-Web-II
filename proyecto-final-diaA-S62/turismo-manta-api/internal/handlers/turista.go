package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type turistaDemo struct {
	ID           int    `json:"id"`
	Nombre       string `json:"nombre"`
	Nacionalidad string `json:"nacionalidad"`
	Idioma       string `json:"idioma"`
}

var turistasHardcoded = []turistaDemo{
	{ID: 1, Nombre: "John Smith", Nacionalidad: "estadounidense", Idioma: "en"},
	{ID: 2, Nombre: "Marie Dubois", Nacionalidad: "francesa", Idioma: "fr"},
	{ID: 3, Nombre: "Hans Müller", Nacionalidad: "alemana", Idioma: "de"},
}

// ListarTuristas atiende GET /api/v1/turistas.
// NO es método de Server, es función suelta. Ese contraste es deliberado.
func ListarTuristas(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(turistasHardcoded); err != nil {
		log.Printf("error codificando JSON: %v", err)
	}
}
