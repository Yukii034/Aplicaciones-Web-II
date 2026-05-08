package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

// Cubre escenarios del método Guardar usando el patrón table-driven.
// Casos cubiertos:
//  1. Caso feliz
//  2. Nombre vacío
//  3. Idioma vacío
//  4. Idioma inválido
//  5. ID duplicado
func TestTuristaMemoria_Guardar(t *testing.T) {

	repo := NewTuristaMemoria()

	// Nuevo turista
	base := models.Turista{
		ID:              1,
		Nombre:          "John Smith",
		Nacionalidad:    "USA",
		IdiomaPreferido: "en",
	}

	if err := repo.Guardar(base); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	casos := []struct {
		nombre    string
		entrada   models.Turista
		esperaErr error
	}{
		{
			nombre: "caso feliz",
			entrada: models.Turista{
				ID:              100,
				Nombre:          "Marie Dubois",
				Nacionalidad:    "Francia",
				IdiomaPreferido: "fr",
			},
			esperaErr: nil,
		},
		{
			nombre: "nombre vacío falla",
			entrada: models.Turista{
				ID:              101,
				Nombre:          "",
				Nacionalidad:    "Italia",
				IdiomaPreferido: "it",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "idioma vacío falla",
			entrada: models.Turista{
				ID:              102,
				Nombre:          "Mario Rossi",
				Nacionalidad:    "Italia",
				IdiomaPreferido: "",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "idioma invalido",
			entrada: models.Turista{
				ID:              102,
				Nombre:          "Mario Rossi",
				Nacionalidad:    "Italia",
				IdiomaPreferido: "ja",
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.Turista{
				ID:              1,
				Nombre:          "Otro Turista",
				Nacionalidad:    "USA",
				IdiomaPreferido: "en",
			},
			esperaErr: errs.ErrYaExiste,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			err := repo.Guardar(c.entrada)

			if !errors.Is(err, c.esperaErr) {
				t.Errorf("Guardar(%q): esperaba %v, obtuvo %v",
					c.entrada.Nombre, c.esperaErr, err)
			}
		})
	}
}

// El test de Buscar por ID cubre:
//  1. Caso feliz
//  2. ID negativo
//  3. ID inexistente
func TestTuristaMemoria_BuscarPorID(t *testing.T) {

	repo := NewTuristaMemoria()

	base := models.Turista{
		ID:              1,
		Nombre:          "John Smith",
		Nacionalidad:    "USA",
		IdiomaPreferido: "en",
	}

	if err := repo.Guardar(base); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	casos := []struct {
		nombre    string
		idBuscar  int
		esperaErr error
	}{
		{
			nombre:    "caso feliz",
			idBuscar:  1,
			esperaErr: nil,
		},
		{
			nombre:    "ID negativo falla",
			idBuscar:  -1,
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre:    "ID inexistente falla",
			idBuscar:  999,
			esperaErr: errs.ErrNoEncontrado,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			_, err := repo.BuscarPorID(c.idBuscar)

			if !errors.Is(err, c.esperaErr) {
				t.Errorf("BuscarPorID(%d): esperaba %v, obtuvo %v",
					c.idBuscar, c.esperaErr, err)
			}
		})
	}
}

// El test Listar verifica que devuelva todos los turistas guardados
func TestTuristaMemoria_Listar(t *testing.T) {

	repo := NewTuristaMemoria()

	//Arrange
	t1 := models.Turista{
		ID:              1,
		Nombre:          "John",
		Nacionalidad:    "USA",
		IdiomaPreferido: "en",
	}

	t2 := models.Turista{
		ID:              2,
		Nombre:          "Mario",
		Nacionalidad:    "Italia",
		IdiomaPreferido: "it",
	}

	_ = repo.Guardar(t1)
	_ = repo.Guardar(t2)

	//Act
	lista := repo.Listar()

	//Assert
	if len(lista) != 2 {
		t.Errorf("esperaba 2 turistas, obtuvo %d", len(lista))
	}
}
