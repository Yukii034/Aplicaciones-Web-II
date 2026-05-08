package storage

import (
	"errors"
	"testing"

	"github.com/uleam/awii/turismo/internal/errs"
	"github.com/uleam/awii/turismo/internal/models"
)

// setupRepos crea los 3 repositorios y siembra un turista y un negocio válidos.
// Esto evita repetir el mismo setup en todos los tests.
func setupRepos(t *testing.T) (*TuristaMemoria, *NegocioMemoria, *CheckInMemoria) {

	turistas := NewTuristaMemoria()
	negocios := NewNegocioMemoria()
	checkins := NewCheckInMemoria(turistas, negocios)

	// Turista nuevo y válido
	err := turistas.Guardar(models.Turista{
		ID:              1,
		Nombre:          "John Smith",
		Nacionalidad:    "USA",
		IdiomaPreferido: "en",
	})

	if err != nil {
		t.Fatalf("setup turista falló: %v", err)
	}

	// Negocio nuevo y válido
	err = negocios.Guardar(models.Negocio{
		ID: 1, Nombre: "Café del Mar",
		Tipo:            "restaurante",
		Ciudad:          "Manta",
		IdiomasHablados: []string{"es", "en"},
		Activo:          true,
	})

	if err != nil {
		t.Fatalf("setup negocio falló: %v", err)
	}

	return turistas, negocios, checkins
}

// El test Guardar cubre los principales escenarios de validación de CheckIn.Guardar
func TestCheckInMemoria_Guardar(t *testing.T) {

	_, _, checkins := setupRepos(t)

	// Nuevo checkin para probar ID duplicado
	base := models.CheckIn{
		ID:           1,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-04-10",
		Calificacion: 5,
	}

	if err := checkins.Guardar(base); err != nil {
		t.Fatalf("setup falló: %v", err)
	}

	casos := []struct {
		nombre    string
		entrada   models.CheckIn
		esperaErr error
	}{
		{
			nombre: "caso feliz",
			entrada: models.CheckIn{
				ID:           100,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-04-11",
				Calificacion: 4,
			},
			esperaErr: nil,
		},
		{
			nombre: "fecha vacía falla",
			entrada: models.CheckIn{
				ID:           101,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "",
				Calificacion: 5,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "calificación 0 falla",
			entrada: models.CheckIn{
				ID:           102,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-04-11",
				Calificacion: 0,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "calificación 6 falla",
			entrada: models.CheckIn{
				ID:           103,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-04-11",
				Calificacion: 6,
			},
			esperaErr: errs.ErrDatosInvalidos,
		},
		{
			nombre: "turista inexistente falla",
			entrada: models.CheckIn{
				ID:           104,
				TuristaID:    999,
				NegocioID:    1,
				Fecha:        "2026-04-11",
				Calificacion: 5,
			},
			esperaErr: errs.ErrNoEncontrado,
		},
		{
			nombre: "negocio inexistente falla",
			entrada: models.CheckIn{
				ID:           105,
				TuristaID:    1,
				NegocioID:    999,
				Fecha:        "2026-04-11",
				Calificacion: 5,
			},
			esperaErr: errs.ErrNoEncontrado,
		},
		{
			nombre: "ID duplicado falla",
			entrada: models.CheckIn{
				ID:           1,
				TuristaID:    1,
				NegocioID:    1,
				Fecha:        "2026-04-11",
				Calificacion: 5,
			},
			esperaErr: errs.ErrYaExiste,
		},
	}

	for _, c := range casos {
		t.Run(c.nombre, func(t *testing.T) {

			err := checkins.Guardar(c.entrada)

			if !errors.Is(err, c.esperaErr) {
				t.Errorf("Guardar(): esperaba %v, obtuvo %v",
					c.esperaErr, err)
			}
		})
	}
}

// El test de buscar verifica:
//  1. Turista con check-ins
//  2. Turista sin check-ins
func TestCheckInMemoria_BuscarPorTurista(t *testing.T) {

	_, _, checkins := setupRepos(t)

	_ = checkins.Guardar(models.CheckIn{
		ID:           1,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-04-10",
		Calificacion: 5,
	})

	_ = checkins.Guardar(models.CheckIn{
		ID:           2,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-04-11",
		Calificacion: 4,
	})

	visitas, err := checkins.BuscarPorTurista(1)

	if err != nil {
		t.Errorf("no esperaba error: %v", err)
	}

	if len(visitas) != 2 {
		t.Errorf("esperaba 2 visitas, obtuvo %d", len(visitas))
	}
}

// El test verifica el comportamiento del método Listar con múltiples check-ins.
func TestCheckInMemoria_Listar(t *testing.T) {

	_, _, checkins := setupRepos(t)

	_ = checkins.Guardar(models.CheckIn{
		ID:           1,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-04-10",
		Calificacion: 5,
	})

	_ = checkins.Guardar(models.CheckIn{
		ID:           2,
		TuristaID:    1,
		NegocioID:    1,
		Fecha:        "2026-04-11",
		Calificacion: 4,
	})

	lista := checkins.Listar()

	if len(lista) != 2 {
		t.Errorf("esperaba 2 check-ins, obtuvo %d", len(lista))
	}
}
